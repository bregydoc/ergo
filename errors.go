package ergo

import (
	"github.com/bregydoc/ergo/schema"
	"github.com/oklog/ulid"
)

// ErrorsBag is a bag for errors transactions
type ErrorsBag interface {
	GetAllErrors() ([]*schema.Error, error)
	GetAllErrorsFromCode(code uint64) ([]*schema.Error, error)
	GetErrorByID(id ulid.ULID) (*schema.Error, error)
	GetErrorByNativeError(errN error) (*schema.Error, error)

	RegisterNewErrorFromNative(errN error, message ...string) (*schema.Error, error)
	RegisterNewError(ergoError *schema.Error) (*schema.Error, error)

	UpdateErrorByID(id ulid.ULID, update *schema.Error) (*schema.Error, error)
	UpdateErrorByNative(errN error, update *schema.Error) (*schema.Error, error)

	RemoveErrorByID(id ulid.ULID) (*schema.Error, error)
	RemoveErrorByNative(errN error) (*schema.Error, error)
}
