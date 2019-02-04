package creators

import (
	"github.com/bregydoc/ergo"
	"github.com/bregydoc/ergo/repositories"
	"golang.org/x/text/language"
)

const defaultImage = "https://pass.idbi.pe/static/img/idpass.png"

// NewDefaultErgo ...
func NewDefaultErgoWithBadger() (*ergo.Ergo, error) {
	temp := "./temp"

	repo, err := repositories.NewBadgerRepo(temp, temp)
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
