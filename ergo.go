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
	AvailableLanguades []language.Tag
	DefaultImage       string
	DefaultActionLink  string
}

// Ergo is a wizard
type Ergo struct {
	opt  *Options
	repo Repository
}

// RegisterNewError implements Wizard intreface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) RegisterNewError(where, explain string, message *UserMessage, withFeedback bool) (*schema.ErrorInstance, error) {
	code := 100 + rand.Int63n(300)

	eType := schema.ErrorType_ONLY_READ
	if withFeedback {
		eType = schema.ErrorType_HUMAN_INTERACTIVE
	}

	creator := &ErrorCreator{
		Where:       where,
		Explain:     explain,
		Image:       ergo.opt.DefaultImage,
		Code:        uint64(code),
		ErrorType:   eType,
		Raw:         explain,
		UserMessage: message,
		Action: &Action{
			Message: "More information",
			Link:    ergo.opt.DefaultActionLink,
		},
	}

	return ergo.repo.SaveNewError(creator)
}

// RegisterFullError implements Wizard intreface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) RegisterFullError(asDev *schema.ErrorDev, asHuman *schema.ErrorHuman, withFeedback bool) (*schema.ErrorInstance, error) {
	panic("unimplemented")
}

// ConsultErrorAsHuman implements Wizard intreface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) ConsultErrorAsHuman(errorID ulid.ULID, languages ...language.Tag) (*schema.ErrorHuman, error) {
	return ergo.repo.GetErrorForHuman(errorID, languages...)
}

// ConsultErrorAsDeveloper implements Wizard intreface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) ConsultErrorAsDeveloper(errorID ulid.ULID) (*schema.ErrorDev, error) {
	return ergo.repo.GetErrorForDev(errorID)
}

// MemorizeNewMessages implements Wizard intreface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) MemorizeNewMessages(errorID ulid.ULID, messages ...*UserMessage) ([]*schema.UserMessage, error) {
	responses := make([]*schema.UserMessage, 0)
	for _, m := range messages {
		resp, err := ergo.repo.SetOneMessageError(errorID, m.Language, m.Message)
		if err != nil {
			return nil, err
		}
		responses = append(responses, resp)
	}

	return responses, nil
}

// ReceiveFeedbackOfUser implements Wizard intreface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) ReceiveFeedbackOfUser(errorID ulid.ULID, feedback *UserFeedback) (*schema.Feedback, error) {
	return ergo.repo.AddFeedbackToUser(errorID, feedback)
}
