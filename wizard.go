package ergo

import (
	"github.com/bregydoc/ergo/schema"
	"github.com/oklog/ulid"
	"golang.org/x/text/language"
)

// PersonType is a kind of people
type PersonType int

// Human is a common people, every people is a human
const Human PersonType = 0

// Dev is a developer, this kind of people can understand more deatils in this reality
const Dev PersonType = 1

// Wizard is a interface can to dialoge with persons and retrive data based on its requirements
type Wizard interface {
	RegisterNewError(where, explain string, message *UserMessage, withFeedback bool) (*schema.ErrorInstance, error)
	RegisterFullError(asDev *schema.ErrorDev, asHuman *schema.ErrorHuman, withFeedback bool) (*schema.ErrorInstance, error)
	ConsultErrorAsHuman(errorID ulid.ULID, languages ...language.Tag) (*schema.ErrorHuman, error)
	ConsultErrorAsDeveloper(errorID ulid.ULID) (*schema.ErrorDev, error)

	// Save new messages jeje
	MemorizeNewMessages(errorID ulid.ULID, messages ...*UserMessage) ([]*schema.UserMessage, error)
	// Save new feedback
	ReceiveFeedbackOfUser(errorID ulid.ULID, feedback *UserFeedback) (*schema.Feedback, error)
}
