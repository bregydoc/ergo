package ergo

import (
	"github.com/bregydoc/ergo/schema"
	"github.com/oklog/ulid"
)

// ErrorsBag is a bag for errors transactions
type Bag interface {
	// Errors
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

	// Languages
	RegisterNewLang(language *schema.Language) (*schema.Language, error)
	AddNewLangToError(errorID ulid.ULID, language *schema.Language, message string) (*schema.Language, error)
	GetLanguageByID(id ulid.ULID) (*schema.Language, error)
	UpdateLanguageByID(id ulid.ULID, update *schema.Language) (*schema.Language, error)
	RemoveLanguageByID(id ulid.ULID) (*schema.Language, error)
}
