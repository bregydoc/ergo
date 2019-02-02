package main

import (
	"encoding/json"
	"fmt"
	"github.com/bregydoc/ergo/creators"
)

func main() {
	e, err := creators.NewDefaultErgoWithBadger()
	if err != nil {
		panic(err)
	}
	// where := devs.TraceError()
	//
	// instance, err := e.RegisterNewError(where,
	// 	http.ErrAbortHandler.Error(),
	// 	&ergo.UserMessage{Message:"Error abort handle", Language:language.English},
	// 	false,
	// )
	//
	// if err != nil {
	// 	panic(err)
	// }

	id := []byte{0x00, 0x00, 0x3b, 0x9a, 0xca, 0x00, 0xa5, 0xe5, 0x15, 0xbc, 0x97, 0xe8, 0x5c, 0xf6, 0x9b, 0xc3}
	forHuman, err := e.ConsultErrorAsHuman(id)
	if err != nil {
		panic(err)
	}

	s, _ := json.MarshalIndent(forHuman, "", "\t")
	fmt.Println(string(s))

	forDev, err := e.ConsultErrorAsDeveloper(id)
	if err != nil {
		panic(err)
	}

	s, _ = json.MarshalIndent(forDev, "", "\t")
	fmt.Println(string(s))

}
