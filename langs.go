package ergo

import (
	"github.com/bregydoc/ergo/schema"
	"github.com/oklog/ulid"
	"golang.org/x/text/language"
)

// Language is a synthetic type
type Language struct {
	schema.Language
}

// NewLanguageByTag generate a new language from tag
func NewLanguageByTag(id ulid.ULID, tag language.Tag) (*schema.Language, error) {
	return &schema.Language{
		Id:   id[:],
		Name: tag.String(),
	}, nil

}

type LanguagesBag interface {
	RegisterNewLang(language *schema.Language) (*schema.Language, error)
	AddNewLangToError(errorID ulid.ULID, language *schema.Language, message string) (*schema.Language, error)
	GetLanguageByID(id ulid.ULID) (*schema.Language, error)
	UpdateLanguageByID(id ulid.ULID, update *schema.Language) (*schema.Language, error)
	RemoveLanguageByID(id ulid.ULID) (*schema.Language, error)
}
