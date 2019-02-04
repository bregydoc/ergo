package creators

import (
	"github.com/bregydoc/ergo"
	"github.com/bregydoc/ergo/repositories"
	"golang.org/x/text/language"
)

const defaultImage = "https://pass.idbi.pe/static/img/idpass.png"

// NewDefaultErgo ...
func NewDefaultErgoWithBadger(dir ...string) (*ergo.Ergo, error) {
	finalDir := "./temp"

	if len(dir) != 0 {
		finalDir = dir[0]
	}

	repo, err := repositories.NewBadgerRepo(finalDir, finalDir)
	if err != nil {
		return nil, err
	}

	o := &ergo.Options{
		DefaultLanguage:    language.English,
		AvailableLanguages: []language.Tag{language.Spanish, language.Chinese, language.Portuguese},
		DefaultImage:       defaultImage,
		DefaultActionLink:  "https://pass.idbi.pe/feedback",
	}

	e := &ergo.Ergo{Opt: o, Repo: repo}

	if err != nil {
		return nil, err
	}

	return e, nil
}
