package main

import (
	"encoding/json"
	"fmt"

	"github.com/bregydoc/ergo"
	"github.com/bregydoc/ergo/creators"
	"github.com/oklog/ulid"
	"golang.org/x/text/language"
)

func mai1n() {
	e, err := creators.NewDefaultErgoWithBadger()
	if err != nil {
		panic(err)
	}

	id := ulid.ULID{0x00, 0x00, 0x3b, 0x9a, 0xca, 0x00, 0xa5, 0xe5, 0x15, 0xbc, 0x97, 0xe8, 0x5c, 0xf6, 0x9b, 0xc3}

	_, err = e.MemorizeNewMessages(id.String(), true,
		&ergo.UserMessage{Language: language.Spanish},
		&ergo.UserMessage{Language: language.Japanese},
		&ergo.UserMessage{Language: language.Korean},
		&ergo.UserMessage{Language: language.Afrikaans, Message: "uga uga"},
	)
	if err != nil {
		panic(err)
	}

	forHuman, err := e.ConsultErrorAsHumanByID(id, language.English, language.Spanish, language.Japanese, language.Korean, language.Afrikaans)
	if err != nil {
		panic(err)
	}

	s, _ := json.MarshalIndent(forHuman, "", "\t")
	fmt.Println(string(s))

	forDev, err := e.ConsultErrorAsDeveloperByID(id)
	if err != nil {
		panic(err)
	}

	s, _ = json.MarshalIndent(forDev, "", "\t")
	fmt.Println(string(s))

}
