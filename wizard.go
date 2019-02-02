package ergo

import (
	"github.com/bregydoc/ergo/schema"
	"golang.org/x/text/language"
)

// PersonType is a kind of people
type PersonType int

// Human is a common people, every people is a human
const Human PersonType = 0

// Dev is a developer, this kind of people can understand more details in this reality
const Dev PersonType = 1

// Wizard is a interface can to dialog with persons and retrieve data based on its requirements
type Wizard interface {
	RegisterNewError(where, explain string, message *UserMessage, withFeedback bool) (*schema.ErrorInstance, error)
	RegisterFullError(asDev *schema.ErrorDev, asHuman *schema.ErrorHuman, withFeedback bool) (*schema.ErrorInstance, error)
	ConsultErrorAsHuman(errorID []byte, languages ...language.Tag) (*schema.ErrorHuman, error)
	ConsultErrorAsDeveloper(errorID []byte) (*schema.ErrorDev, error)

	// Save new messages
	MemorizeNewMessages(errorID []byte, messages ...*UserMessage) ([]*schema.UserMessage, error)
	// Save new feedback
	ReceiveFeedbackOfUser(errorID []byte, feedback *UserFeedback) (*schema.Feedback, error)
}
