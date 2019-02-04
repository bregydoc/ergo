package ergo

import (
	"github.com/bregydoc/ergo/schema"
	"github.com/oklog/ulid"
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
	ByID    ulid.ULID
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
	RegisterNewUserMessage(errorID ulid.ULID, uMessage *UserMessage) (*schema.UserMessage, error)
	AddFeedbackToUser(errorID ulid.ULID, feedback *UserFeedback) (*schema.Feedback, error)

	GetErrorInstance(errorID ulid.ULID) (*schema.ErrorInstance, error)
	GetErrorForHuman(errorID ulid.ULID, languages ...language.Tag) (*schema.ErrorHuman, error)
	GetErrorForDev(errorID ulid.ULID) (*schema.ErrorDev, error)

	// UpdateError returns a ErrorDev cause I understand who call this method is a dev
	UpdateError(errorID ulid.ULID, update *ErrorUpdate) (*schema.ErrorDev, error)
	DeleteError(errorID ulid.ULID) error

	// Temporal
	SetOneMessageError(errorID ulid.ULID, language language.Tag, message string) (*schema.UserMessage, error)
}
