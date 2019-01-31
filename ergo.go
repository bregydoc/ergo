package ergo

import (
	"github.com/bregydoc/ergo/schema"
	"github.com/oklog/ulid"
)

// Ergo is a wizard
type Ergo struct {
	repo Repository
}

// RegisterNewError implements Wizard intreface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) RegisterNewError(where, cause string, message *UserMessage, withFeedback bool) (*schema.ErrorInstance, error) {
	return nil, nil
}

// RegisterFullError implements Wizard intreface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) RegisterFullError(asDev *schema.ErrorDev, asHuman *schema.ErrorHuman, withFeedback bool) (*schema.ErrorInstance, error) {
	return nil, nil
}

// ConsultErrorAsHuman implements Wizard intreface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) ConsultErrorAsHuman(errorID ulid.ULID) (*schema.ErrorHuman, error) {
	return nil, nil
}

// ConsultErrorAsDeveloper implements Wizard intreface, then Ergo is a wizard, a wizard of your errors
func (ergo *Ergo) ConsultErrorAsDeveloper(errorID ulid.ULID) (*schema.ErrorDev, error) {
	return nil, nil
}
