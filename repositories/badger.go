package repositories

import (
	"encoding/binary"

	"github.com/bregydoc/ergo"
	"github.com/bregydoc/ergo/schema"
	"github.com/dgraph-io/badger"
	"github.com/gogo/protobuf/proto"
	"golang.org/x/text/language"
)

var instancePrefix = []byte("instance")
var devPrefix = []byte("dev")
var humanPrefix = []byte("human")

var languagePrefix = []byte("lang")

func getErrorInstanceID(id string) []byte {
	return append(instancePrefix, []byte(id)...)
}

func getErrorDevID(id string) []byte {
	return append(devPrefix, []byte(id)...)
}

func getErrorHumanID(id string) []byte {
	return append(humanPrefix, []byte(id)...)
}

func getErrorMessageLanguageID(errorID string, lang language.Tag) []byte {
	c := append([]byte(errorID), []byte(lang.String())...)
	return append(languagePrefix, c...)
}

// BadgerRepo implement ergo repo
type BadgerRepo struct {
	db                     *badger.DB
	onNewErrorHasBeenSaved *func(value *schema.ErrorInstance)
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

	newErrorID := ergo.UlidGen.New().String()

	// --------- --------- ---------
	// On early step I'm going to register the code:ulid pair in the database
	codeBin := make([]byte, 8)
	binary.LittleEndian.PutUint64(codeBin, seed.Code)

	err := txn.Set(codeBin, []byte(newErrorID))
	if err != nil {
		return nil, err
	}
	// --------- --------- ---------

	// First, create the instance
	instance := &schema.ErrorInstance{
		Id:   newErrorID,
		Code: seed.Code,
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
		Id:       newErrorID,
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
		Messages: []*schema.UserMessage{},
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
	// If the error was registered successfully
	// launch callback
	go func() {
		if b.onNewErrorHasBeenSaved != nil {
			callback := *b.onNewErrorHasBeenSaved
			callback(instance)
		}
	}()

	return instance, nil
}

// RegisterNewUserMessage implements the Bag interface of Ergo
func (b *BadgerRepo) RegisterNewUserMessage(errorID string, uMessage *ergo.UserMessage) (*schema.UserMessage, error) {
	txn := b.db.NewTransaction(true)
	defer txn.Discard()

	mID := getErrorMessageLanguageID(errorID, uMessage.Language)
	m := &schema.UserMessage{
		Id:       errorID,
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
func (b *BadgerRepo) AddFeedbackToUser(errorID string, feedback *ergo.UserFeedback) (*schema.Feedback, error) {
	txn := b.db.NewTransaction(true)
	defer txn.Discard()

	item, err := txn.Get(getErrorDevID(errorID))
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

	err = txn.Set(getErrorDevID(errorID), dataError)
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
func (b *BadgerRepo) GetErrorInstance(errorID string) (*schema.ErrorInstance, error) {
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

// GetErrorInstanceByCode implements the Bag interface of Ergo
func (b *BadgerRepo) GetErrorInstanceByCode(code uint64) (*schema.ErrorInstance, error) {
	txn := b.db.NewTransaction(false)
	defer txn.Discard()

	codeBin := make([]byte, 8)
	binary.LittleEndian.PutUint64(codeBin, code)

	item, err := txn.Get(codeBin)
	if err != nil {
		return nil, err
	}

	data, err := item.ValueCopy(nil)
	// data may will be the stringify ulid

	return b.GetErrorInstance(string(data))

}

// GetErrorForHuman implements the Bag interface of Ergo
func (b *BadgerRepo) GetErrorForHuman(errorID string, languages ...language.Tag) (*schema.ErrorHuman, error) {
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
func (b *BadgerRepo) GetErrorForDev(errorID string) (*schema.ErrorDev, error) {
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

// GetAllRegisteredErrors implements the Bag interface of Ergo
func (b *BadgerRepo) GetAllRegisteredErrors() ([]*schema.ErrorInstance, error) {
	txn := b.db.NewTransaction(false)
	defer txn.Discard()

	iter := txn.NewIterator(badger.IteratorOptions{
		Prefix:         instancePrefix,
		PrefetchValues: true,
	})

	defer iter.Close()

	instances := make([]*schema.ErrorInstance, 0)
	for iter.Rewind(); iter.Valid(); iter.Next() {
		item := iter.Item()
		data, err := item.ValueCopy(nil)
		if err != nil {
			return nil, err
		}

		var inst schema.ErrorInstance

		err = proto.Unmarshal(data, &inst)
		if err != nil {
			return nil, err
		}

		instances = append(instances, &inst)
	}

	return instances, nil
}

// GetAllErrorsForDev implements the Bag interface of Ergo
func (b *BadgerRepo) GetAllErrorsForDev() ([]*schema.ErrorDev, error) {
	txn := b.db.NewTransaction(false)
	defer txn.Discard()

	iter := txn.NewIterator(badger.IteratorOptions{
		Prefix:         devPrefix,
		PrefetchValues: true,
	})

	defer iter.Close()

	errorsForDev := make([]*schema.ErrorDev, 0)
	for iter.Rewind(); iter.Valid(); iter.Next() {
		item := iter.Item()
		data, err := item.ValueCopy(nil)
		if err != nil {
			return nil, err
		}

		var dev schema.ErrorDev

		err = proto.Unmarshal(data, &dev)
		if err != nil {
			return nil, err
		}

		errorsForDev = append(errorsForDev, &dev)
	}

	return errorsForDev, nil
}

// GetAllErrorsForUI implements the Bag interface of Ergo
func (b *BadgerRepo) GetAllErrorsForUI() ([]*ergo.ErrorSummary, error) {
	txn := b.db.NewTransaction(false)
	defer txn.Discard()

	iter := txn.NewIterator(badger.IteratorOptions{
		Prefix:         instancePrefix,
		PrefetchValues: true,
	})

	defer iter.Close()

	allErrors := make([]*ergo.ErrorSummary, 0)
	for iter.Rewind(); iter.Valid(); iter.Next() {
		instanceItem := iter.Item()
		data, err := instanceItem.ValueCopy(nil)
		if err != nil {
			return nil, err
		}
		instance := new(schema.ErrorInstance)
		err = proto.Unmarshal(data, instance)
		if err != nil {
			return nil, err
		}

		forHuman, err := b.GetErrorForHuman(instance.Id, language.English)
		if err != nil {
			return nil, err
		}

		forDev, err := b.GetErrorForDev(instance.Id)
		if err != nil {
			return nil, err
		}
		action := &schema.Action{
			Link:    "",
			Message: "",
		}

		if len(forHuman.Action) > 0 {
			action = forHuman.Action[0]
		}

		message := ""
		if len(forHuman.Messages) > 0 {
			message = forHuman.Messages[0].Message
		}

		e := &ergo.ErrorSummary{
			ID:                 instance.Id,
			Code:               instance.Code,
			Explain:            forDev.Explain,
			Image:              forHuman.Image,
			ActionLink:         action.Link,
			ActionMessage:      action.Message,
			EnglishUserMessage: message,
			Type:               forDev.Type.String(),
		}

		allErrors = append(allErrors, e)
	}

	return allErrors, nil
}

// UpdateError implements the Bag interface of Ergo
func (b *BadgerRepo) UpdateError(errorID string, update *ergo.ErrorUpdate) (*schema.ErrorDev, error) {
	panic("unimplemented")
	// return nil, nil
}

// DeleteError implements the Bag interface of Ergo
func (b *BadgerRepo) DeleteError(errorID string) error {
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
func (b *BadgerRepo) SetOneMessageError(errorID string, language language.Tag, message string) (*schema.UserMessage, error) {
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

// OnNewErrorHasBeenSaved implements the Bag interface of Ergo
func (b *BadgerRepo) OnNewErrorHasBeenSaved(callback func(value *schema.ErrorInstance)) error {
	b.onNewErrorHasBeenSaved = &callback
	return nil
}
