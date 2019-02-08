package ergo

import (
	"github.com/bregydoc/ergo/schema"
	"golang.org/x/text/language"
)

// UserMessage is a util struct for create first message error
type UserMessage struct {
	Language language.Tag
	Message  string
}

// Action describe a possible actions for users
type Action struct {
	Link    string
	Message string
}

// ErrorCreator is a packed struct util for create a new error
type ErrorCreator struct {
	ErrorType   schema.ErrorType
	Code        uint64
	Explain     string
	Where       string
	Raw         string
	UserMessage *UserMessage
	Action      *Action
	Image       string
	SuggestedID []byte
}

// UserFeedback is a util struct to create user feedback response
type UserFeedback struct {
	By      string
	ByID    string
	Message string
}

// ErrorUpdate describes the update payload
type ErrorUpdate struct {
	ErrorType schema.ErrorType
	Explain   string
	Where     string
	Raw       string
	Image     string
	Actions   []Action
}

// Repository is a bag for errors
type Repository interface {
	SaveNewError(seed *ErrorCreator) (*schema.ErrorInstance, error)
	RegisterNewUserMessage(errorID string, uMessage *UserMessage) (*schema.UserMessage, error)
	AddFeedbackToUser(errorID string, feedback *UserFeedback) (*schema.Feedback, error)

	GetErrorInstance(errorID string) (*schema.ErrorInstance, error)
	GetErrorInstanceByCode(code uint64) (*schema.ErrorInstance, error)
	GetErrorForHuman(errorID string, languages ...language.Tag) (*schema.ErrorHuman, error)
	GetErrorForDev(errorID string) (*schema.ErrorDev, error)

	GetAllRegisteredErrors() ([]*schema.ErrorInstance, error)
	GetAllErrorsForDev() ([]*schema.ErrorDev, error)
	GetAllErrorsForUI() ([]*ErrorSummary, error)

	// UpdateError returns a ErrorDev cause I understand who call this method is a dev
	UpdateError(errorID string, update *ErrorUpdate) (*schema.ErrorDev, error)
	DeleteError(errorID string) error

	// Temporal
	SetOneMessageError(errorID string, language language.Tag, message string) (*schema.UserMessage, error)

	//Synthetic events
	OnNewErrorHasBeenSaved(callback func(value *schema.ErrorInstance)) error
}
