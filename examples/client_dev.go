package main

import (
	"context"
	"flag"

	"github.com/k0kubun/pp"

	"github.com/bregydoc/ergo/devs"
	"github.com/bregydoc/ergo/schema"
	ergocon "github.com/bregydoc/ergo/service"
	"google.golang.org/grpc"
)

func main() {
	code := flag.Int64("code", 133, "Define the code of your error")
	flag.Parse()

	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := ergocon.NewErgoClient(conn)

	resp, err := client.RegisterNewError(context.TODO(), &schema.ErrorSeed{
		Code:            uint64(*code),
		Explain:         "My first custom ergo error",
		MessageLanguage: "en",
		MessageContent:  "Hello world, I'm an error user message",
		Where:           devs.TraceError(),
	})
	if err != nil {
		panic(err)
	}

	pp.Println(resp)
}
