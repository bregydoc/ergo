package main

import (
	"flag"
	"fmt"
	"github.com/bregydoc/ergo"
	"github.com/bregydoc/ergo/creators"
	"github.com/bregydoc/ergo/schema"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := flag.Int64("port", 10000, "Define the service port")
	dir := flag.String("dir", "./ergo_temp", "Workspace for the db")
	flag.Parse()

	e, err := creators.NewDefaultErgoWithBadger(*dir)
	if err != nil {
		panic(err)
	}
	server := ergo.NewErgoServer(e)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	schema.RegisterErgoServer(grpcServer, server)
	log.Printf("listening on :%d\n", *port)

	err = grpcServer.Serve(lis)
}
