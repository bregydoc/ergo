package ergo

import (
	"math/rand"

	"github.com/bregydoc/ergo/schema"
	"github.com/oklog/ulid"
	"golang.org/x/text/language"
)

// Options is a struct to configure to wizard
type Options struct {
	DefaultLanguage    language.Tag
	AvailableLanguages []language.Tag
	DefaultImage       string
	DefaultActionLink  string
}

// Ergo is a wizard
type Ergo struct {
	Opt  *Options
	Repo Repository
}

// RegisterNewError implements Wizard interface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) RegisterNewError(where, explain string, message *UserMessage, withFeedback bool, suggestedID ...[]byte) (*schema.ErrorInstance, error) {
	var finalCode = uint64(100 + rand.Int63n(300))

	var sugID []byte
	if len(suggestedID) != 0 {
		sugID = suggestedID[0]
	} else {
		sugID = nil
	}

	eType := schema.ErrorType_ONLY_READ
	if withFeedback {
		eType = schema.ErrorType_HUMAN_INTERACTIVE
	}

	creator := &ErrorCreator{
		Where:       where,
		Explain:     explain,
		Image:       ergo.Opt.DefaultImage,
		Code:        finalCode,
		ErrorType:   eType,
		Raw:         explain,
		UserMessage: message,
		Action: &Action{
			Message: "More information",
			Link:    ergo.Opt.DefaultActionLink,
		},
		SuggestedID: sugID,
	}

	return ergo.Repo.SaveNewError(creator)
}

// RegisterFullError implements Wizard interface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) RegisterFullError(asDev *schema.ErrorDev, asHuman *schema.ErrorHuman, withFeedback bool) (*schema.ErrorInstance, error) {
	panic("unimplemented")
}

// ConsultErrorAsHuman implements Wizard interface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) ConsultErrorAsHuman(errorID []byte, languages ...language.Tag) (*schema.ErrorHuman, error) {
	var id ulid.ULID
	copy(id[:], errorID)
	return ergo.Repo.GetErrorForHuman(id, languages...)
}

// ConsultErrorAsDeveloper implements Wizard interface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) ConsultErrorAsDeveloper(errorID []byte) (*schema.ErrorDev, error) {
	var id ulid.ULID
	copy(id[:], errorID)
	return ergo.Repo.GetErrorForDev(id)
}

// MemorizeNewMessages implements Wizard interface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) MemorizeNewMessages(errorID []byte, messages ...*UserMessage) ([]*schema.UserMessage, error) {
	var id ulid.ULID
	copy(id[:], errorID)
	responses := make([]*schema.UserMessage, 0)
	for _, m := range messages {
		resp, err := ergo.Repo.SetOneMessageError(id, m.Language, m.Message)
		if err != nil {
			return nil, err
		}
		responses = append(responses, resp)
	}

	return responses, nil
}

// ReceiveFeedbackOfUser implements Wizard interface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) ReceiveFeedbackOfUser(errorID []byte, feedback *UserFeedback) (*schema.Feedback, error) {
	var id ulid.ULID
	copy(id[:], errorID)
	return ergo.Repo.AddFeedbackToUser(id, feedback)
}
