package ergo

import (
	"github.com/bregydoc/ergo/schema"
	"github.com/oklog/ulid"
)

// PersonType is a kind of people
type PersonType int

// Human is a common people, every people is a human
const Human PersonType = 0

// Dev is a developer, this kind of people can understand more deatils in this reality
const Dev PersonType = 1

// Wizard is a interface can to dialoge with persons and retrive data based on its requirements
type Wizard interface {
	ConsultError(as PersonType, errorID ulid.ULID) (*schema.ErrorInstance, error)
}
