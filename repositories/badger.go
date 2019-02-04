package repositories

import (
	"github.com/bregydoc/ergo"
	"github.com/bregydoc/ergo/schema"
	"github.com/dgraph-io/badger"
	"github.com/gogo/protobuf/proto"
	"github.com/oklog/ulid"
	"golang.org/x/text/language"
)

var instanceSuffix = []byte("instance")
var devSuffix = []byte("dev")
var humanSuffix = []byte("human")

var languagePrefix = []byte("lang")

func getErrorInstanceID(id ulid.ULID) []byte {
	return append(id[:], instanceSuffix...)
}

func getErrorDevID(id ulid.ULID) []byte {
	return append(id[:], devSuffix...)
}

func getErrorHumanID(id ulid.ULID) []byte {
	return append(id[:], humanSuffix...)
}

func getErrorMessageLanguageID(errorID ulid.ULID, lang language.Tag) []byte {
	c := append(errorID[:], []byte(lang.String())...)
	return append(languagePrefix, c...)
}

// BadgerRepo implement ergo repo
type BadgerRepo struct {
	db *badger.DB
}

// NewBadgerRepo returns a new repo with Badger as a backend
func NewBadgerRepo(dir, valueDir string, opts ...badger.Options) (*BadgerRepo, error) {
	// db, err := badger.Open(opts)
	opt := badger.DefaultOptions
	if len(opts) > 0 {
		opt = opts[1]
	}
	opt.Dir = dir
	opt.ValueDir = valueDir
	db, err := badger.Open(opt)
	if err != nil {
		return nil, err
	}

	return &BadgerRepo{
		db: db,
	}, nil
}

// SaveNewError implements the Bag interface of Ergo
func (b *BadgerRepo) SaveNewError(seed *ergo.ErrorCreator) (*schema.ErrorInstance, error) {
	txn := b.db.NewTransaction(true)
	defer txn.Discard()

	var newErrorID ulid.ULID
	if seed.SuggestedID != nil {
		var err error
		newErrorID, err = ulid.Parse(string(seed.SuggestedID[0]))
		if err != nil {
			return nil, err
		}
	} else {
		newErrorID = ergo.UlidGen.New()
	}

	// First, create the instance
	instance := &schema.ErrorInstance{
		Id:   newErrorID[:],
		Type: seed.ErrorType,
	}

	dataInstance, err := proto.Marshal(instance)
	if err != nil {
		return nil, err
	}

	err = txn.Set(getErrorInstanceID(newErrorID), dataInstance)
	if err != nil {
		return nil, err
	}

	// Second, generate the dev error
	dev := &schema.ErrorDev{
		Id:       newErrorID[:],
		Type:     seed.ErrorType,
		Code:     seed.Code,
		Explain:  seed.Explain,
		Feedback: []*schema.Feedback{},
		Raw:      seed.Raw,
		Where:    seed.Where,
	}

	dataDev, err := proto.Marshal(dev)
	if err != nil {
		return nil, err
	}

	err = txn.Set(getErrorDevID(newErrorID), dataDev)
	if err != nil {
		return nil, err
	}

	// Finally, generating the human error

	// Saving the first user message
	uMessageID := getErrorMessageLanguageID(newErrorID, seed.UserMessage.Language)
	uMessage := &schema.UserMessage{
		Id:       newErrorID[:],
		Language: seed.UserMessage.Language.String(),
		Message:  seed.UserMessage.Message,
	}
	dataUMessage, err := proto.Marshal(uMessage)
	if err != nil {
		return nil, err
	}

	err = txn.Set(uMessageID, dataUMessage)
	if err != nil {
		return nil, err
	}

	// Creating human error with user message created
	human := &schema.ErrorHuman{
		Id:       newErrorID[:],
		Type:     seed.ErrorType,
		Image:    seed.Image,
		Messages: []*schema.UserMessage{uMessage},
		Action:   []*schema.Action{{Link: seed.Action.Link, Message: seed.Action.Message}},
	}

	dataHuman, err := proto.Marshal(human)
	if err != nil {
		return nil, err
	}

	err = txn.Set(getErrorHumanID(newErrorID), dataHuman)
	if err != nil {
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return instance, nil
}

// RegisterNewUserMessage implements the Bag interface of Ergo
func (b *BadgerRepo) RegisterNewUserMessage(errorID ulid.ULID, uMessage *ergo.UserMessage) (*schema.UserMessage, error) {
	txn := b.db.NewTransaction(true)
	defer txn.Discard()

	mID := getErrorMessageLanguageID(errorID, uMessage.Language)
	m := &schema.UserMessage{
		Id:       errorID[:],
		Language: uMessage.Language.String(),
		Message:  uMessage.Message,
	}

	mData, err := proto.Marshal(m)
	if err != nil {
		return nil, err
	}

	err = txn.Set(mID, mData)
	if err != nil {
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return m, nil
}

// AddFeedbackToUser implements the Bag interface of Ergo
func (b *BadgerRepo) AddFeedbackToUser(errorID ulid.ULID, feedback *ergo.UserFeedback) (*schema.Feedback, error) {
	txn := b.db.NewTransaction(true)
	defer txn.Discard()

	item, err := txn.Get(errorID[:])
	if err != nil {
		return nil, err
	}

	dataError, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}

	// ** Marshal-Unmarshal process **
	errorDev := new(schema.ErrorDev)
	err = proto.Unmarshal(dataError, errorDev)
	if err != nil {
		return nil, err
	}

	if errorDev.Feedback == nil {
		errorDev.Feedback = []*schema.Feedback{}
	}

	errorDev.Feedback = append(errorDev.Feedback, &schema.Feedback{
		By:      feedback.By,
		ByID:    feedback.ByID[:],
		Message: feedback.Message,
	})

	dataError, err = proto.Marshal(errorDev)
	if err != nil {
		return nil, err
	}
	// ** Marshal-Unmarshal process **

	err = txn.Set(errorID[:], dataError)
	if err != nil {
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// GetErrorInstance implements the Bag interface of Ergo
func (b *BadgerRepo) GetErrorInstance(errorID ulid.ULID) (*schema.ErrorInstance, error) {
	txn := b.db.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get(getErrorInstanceID(errorID))
	if err != nil {
		return nil, err
	}

	errorData, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}
	errorInstance := new(schema.ErrorInstance)

	err = proto.Unmarshal(errorData, errorInstance)
	if err != nil {
		return nil, err
	}

	return errorInstance, nil
}

// GetErrorForHuman implements the Bag interface of Ergo
func (b *BadgerRepo) GetErrorForHuman(errorID ulid.ULID, languages ...language.Tag) (*schema.ErrorHuman, error) {
	txn := b.db.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get(getErrorHumanID(errorID))
	if err != nil {
		return nil, err
	}

	errorData, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}
	errorHuman := new(schema.ErrorHuman)

	err = proto.Unmarshal(errorData, errorHuman)
	if err != nil {
		return nil, err
	}

	if errorHuman.Messages == nil {
		errorHuman.Messages = []*schema.UserMessage{}
	}
	// Filling messages
	for _, l := range languages {
		langID := getErrorMessageLanguageID(errorID, l)
		item, err := txn.Get(langID)
		if err != nil {
			if err != badger.ErrKeyNotFound {
				return nil, err
			}
			continue
		}

		languageData, err := item.ValueCopy(nil)
		if err != nil {
			return nil, err
		}

		message := new(schema.UserMessage)
		err = proto.Unmarshal(languageData, message)
		if err != nil {
			return nil, err
		}

		errorHuman.Messages = append(errorHuman.Messages, message)
	}
	// Filling end

	return errorHuman, nil
}

// GetErrorForDev implements the Bag interface of Ergo
func (b *BadgerRepo) GetErrorForDev(errorID ulid.ULID) (*schema.ErrorDev, error) {
	txn := b.db.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get(getErrorDevID(errorID))
	if err != nil {
		return nil, err
	}

	errorData, err := item.ValueCopy(nil)
	if err != nil {
		return nil, err
	}
	errorDev := new(schema.ErrorDev)

	err = proto.Unmarshal(errorData, errorDev)
	if err != nil {
		return nil, err
	}

	return errorDev, nil
}

// UpdateError implements the Bag interface of Ergo
func (b *BadgerRepo) UpdateError(errorID ulid.ULID, update *ergo.ErrorUpdate) (*schema.ErrorDev, error) {
	panic("unimplemented")
	// return nil, nil
}

// DeleteError implements the Bag interface of Ergo
func (b *BadgerRepo) DeleteError(errorID ulid.ULID) error {
	txn := b.db.NewTransaction(true)
	defer txn.Discard()

	// Instance
	err := txn.Delete(getErrorInstanceID(errorID))
	if err != nil {
		return err
	}

	// Dev
	err = txn.Delete(getErrorDevID(errorID))
	if err != nil {
		return err
	}

	// Human
	err = txn.Delete(getErrorHumanID(errorID))
	if err != nil {
		return err
	}

	err = txn.Commit()
	if err != nil {
		return err
	}

	return nil
}

// SetOneMessageError implements the Bag interface of Ergo
func (b *BadgerRepo) SetOneMessageError(errorID ulid.ULID, language language.Tag, message string) (*schema.UserMessage, error) {
	txn := b.db.NewTransaction(true)
	defer txn.Discard()

	mID := getErrorMessageLanguageID(errorID, language)
	m := &schema.UserMessage{
		Id:       errorID[:],
		Language: language.String(),
		Message:  message,
	}

	mData, err := proto.Marshal(m)
	if err != nil {
		return nil, err
	}

	err = txn.Set(mID, mData)
	if err != nil {
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return m, nil
}
