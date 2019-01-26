package ergo

import (
	"github.com/bregydoc/ergo/schema"
	"github.com/oklog/ulid"
	"golang.org/x/text/language"
)

var DefaultLanguage = language.English

type Service interface {
	RegisterErrorFromNative(errN error, extraLanguages ...map[language.Tag]string) (*schema.Error, error)
	RegisterError(errMsg string, defaultLanguage language.Tag, extraLanguages ...map[language.Tag]string) (*schema.Error, error)

	UpdateErrorByID(id ulid.ULID, ergoError schema.Error) (*schema.Error, error)

	GetErrorByID(id ulid.ULID) (*schema.Error, error)
	GetErrorMessageByLanguage(errorID ulid.ULID, lang language.Tag, withDefault ...bool) (string, error)
	GetDefaultErrorMessage(errorID ulid.ULID, defaultMessage ...string) (string, error)
}
