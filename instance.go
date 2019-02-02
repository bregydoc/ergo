package ergo

import (
	"golang.org/x/text/language"
)

const defaultImage = "https://pass.idbi.pe/static/img/idpass.png"

// NewErgo returns a new ergo instance service struct
func NewErgo(repo Repository, opts ...Options) (*Ergo, error) {

	o := &Options{
		DefaultLanguage:    language.English,
		AvailableLanguages: []language.Tag{language.Spanish, language.Chinese, language.Portuguese},
		DefaultImage:       defaultImage,
		DefaultActionLink:  "https://pass.idbi.pe/feedback",
	}

	if len(opts) > 0 {
		o = &opts[0]
	}

	return &Ergo{Opt: o, Repo: repo}, nil
}
