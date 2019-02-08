package ergo

import (
	"math/rand"
	"strings"

	"github.com/bregydoc/gtranslate"

	"github.com/bregydoc/ergo/schema"
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
func (ergo *Ergo) RegisterNewError(where, explain string, message *UserMessage, withFeedback bool) (*schema.ErrorInstance, error) {
	var finalCode = uint64(100 + rand.Int63n(300))

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
	}

	instError, err := ergo.Repo.SaveNewError(creator)
	if err != nil {
		return nil, err
	}

	// Verify if message language is English
	if message.Language.String() != language.English.String() { // If isn't
		// Creating english user message for this error
		engMessage, err := gtranslate.TranslateWithFromTo(message.Message, gtranslate.FromTo{From: "auto", To: "en"})
		if err != nil {
			engMessage = ""
		}
		ergo.Repo.SetOneMessageError(instError.Id, language.English, engMessage)

	}

	return instError, nil
}

// RegisterNewErrorWithCode implements Wizard interface
func (ergo *Ergo) RegisterNewErrorWithCode(where, explain string, code uint64, message *UserMessage, withFeedback bool) (*schema.ErrorInstance, error) {
	eType := schema.ErrorType_ONLY_READ
	if withFeedback {
		eType = schema.ErrorType_HUMAN_INTERACTIVE
	}

	creator := &ErrorCreator{
		Where:       where,
		Explain:     explain,
		Image:       ergo.Opt.DefaultImage,
		Code:        code,
		ErrorType:   eType,
		Raw:         explain,
		UserMessage: message,
		Action: &Action{
			Message: "More information",
			Link:    ergo.Opt.DefaultActionLink,
		},
	}

	instError, err := ergo.Repo.SaveNewError(creator)
	if err != nil {
		return nil, err
	}

	// Verify if message language is English
	if message.Language.String() != language.English.String() { // If isn't
		// Creating english user message for this error
		engMessage, err := gtranslate.TranslateWithFromTo(message.Message, gtranslate.FromTo{From: "auto", To: "en"})
		if err != nil {
			engMessage = ""
		}
		ergo.Repo.SetOneMessageError(instError.Id, language.English, engMessage)

	}

	return instError, nil
}

// RegisterFullError implements Wizard interface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) RegisterFullError(asDev *schema.ErrorDev, asHuman *schema.ErrorHuman, withFeedback bool) (*schema.ErrorInstance, error) {
	panic("unimplemented")
}

// ConsultErrorAsHumanByID implements Wizard interface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) ConsultErrorAsHumanByID(errorID string, languages ...language.Tag) (*schema.ErrorHuman, error) {
	forHuman, err := ergo.Repo.GetErrorForHuman(errorID, languages...)
	if err != nil {
		if strings.Contains(err.Error(), "found") {
			return unknownErrorForHumans, nil
		}
		return nil, err
	}

	return forHuman, nil
}

// ConsultErrorAsDeveloperByID implements Wizard interface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) ConsultErrorAsDeveloperByID(errorID string) (*schema.ErrorDev, error) {
	forDev, err := ergo.Repo.GetErrorForDev(errorID)
	if err != nil {
		if strings.Contains(err.Error(), "found") {
			return unknownErrorForDevelopers, nil
		}
		return nil, err
	}

	return forDev, nil
}

// ConsultErrorAsHumanByCode implements Wizard interface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) ConsultErrorAsHumanByCode(code uint64, languages ...language.Tag) (*schema.ErrorHuman, error) {
	errorInstance, err := ergo.Repo.GetErrorInstanceByCode(code)
	if err != nil {
		return nil, err
	}

	return ergo.ConsultErrorAsHumanByID(errorInstance.Id, languages...)
}

// ConsultErrorAsDeveloperByCode implements Wizard interface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) ConsultErrorAsDeveloperByCode(code uint64) (*schema.ErrorDev, error) {
	errorInstance, err := ergo.Repo.GetErrorInstanceByCode(code)
	if err != nil {
		return nil, err
	}
	return ergo.ConsultErrorAsDeveloperByID(errorInstance.Id)
}

// MemorizeNewMessages implements Wizard interface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) MemorizeNewMessages(errorID string, withAutoTranslate bool, messages ...*UserMessage) ([]*schema.UserMessage, error) {
	responses := make([]*schema.UserMessage, 0)
	var ergoError *schema.ErrorHuman

	for _, m := range messages {

		// If the message not exist, we can auto translate it
		message := m.Message
		if withAutoTranslate {
			if m.Message == "" {
				if ergoError == nil {
					var err error
					ergoError, err = ergo.ConsultErrorAsHumanByID(errorID, language.English)
					if err != nil {
						return nil, err
					}
				}

				if len(ergoError.Messages) == 0 {
					continue
				}

				// I expected a english message
				inEnglish := ergoError.Messages[0].Message
				var err error
				message, err = gtranslate.Translate(inEnglish, language.English, m.Language)
				if err != nil {
					continue
				}

			}
		}

		resp, err := ergo.Repo.SetOneMessageError(errorID, m.Language, message)
		if err != nil {
			return nil, err
		}
		responses = append(responses, resp)
	}

	return responses, nil
}

// ReceiveFeedbackOfUser implements Wizard interface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) ReceiveFeedbackOfUser(errorID string, feedback *UserFeedback) (*schema.Feedback, error) {
	return ergo.Repo.AddFeedbackToUser(errorID, feedback)
}
