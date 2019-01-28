package ergo

import (
	"github.com/bregydoc/ergo/schema"
	"github.com/oklog/ulid"
	"golang.org/x/text/language"
)

// NewLanguageByTag generate a new language from tag
func NewLanguageByTag(id ulid.ULID, tag language.Tag) (*schema.Language, error) {
	return &schema.Language{
		Id:   id[:],
		Name: tag.String(),
	}, nil

}
