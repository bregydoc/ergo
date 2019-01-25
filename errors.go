package ergo

import "github.com/oklog/ulid"

// ErrorsBag is a bag for errors transactions
type ErrorsBag interface {
	GetAllErrors() ([]*Error, error)
	GetAllErrorsFromCode(code uint64) ([]*Error, error)
	GetErrorByID(id ulid.ULID) (*Error, error)
	GetErrorByNativeError(err error) (*Error, error)

	RegisterNewErrorFromNative(err error, message ...string) (*Error, error)
	RegisterNewError(err *Error) (*Error, error)

	RemoveErrorByID(id ulid.ULID) (*Error, error)
	RemoveErrorByNative(err error) (*Error, error)
}
