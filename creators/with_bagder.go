package creators

import (
	"github.com/bregydoc/ergo"
	"github.com/bregydoc/ergo/repositories"
	"github.com/oklog/ulid"
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

	// This error is unknown error, util for the first error in ergo
	unknownID := ulid.MustParse("0000000000000000")
	_, err = e.RegisterNewError(
		"ergo",
		"unknown error, ergo could'nt found",
		&ergo.UserMessage{
			Language: language.English,
			Message:  "Sorry, we not know this error, you can send a feedback",
		},
		true,
		unknownID[:], // Null ulid is for unknown error
	)

	if err != nil {
		return nil, err
	}

	return e, nil
}
