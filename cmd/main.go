package main

import (
	"flag"
	"fmt"
	"github.com/bregydoc/ergo"
	"github.com/bregydoc/ergo/creators"
	"github.com/bregydoc/ergo/schema"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	devPort := flag.Int64("dev-port", 10000, "Define the service dev port")
	humanPort := flag.Int64("human-port", 4200, "Define the human server port")
	dir := flag.String("dir", "./ergo_temp", "Workspace for the db")
	flag.Parse()

	e, err := creators.NewDefaultErgoWithBadger(*dir)
	if err != nil {
		panic(err)
	}
	server := ergo.NewErgoServer(e)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *devPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	schema.RegisterErgoServer(grpcServer, server)

	go func() {
		log.Printf("[For Developers] listening on :%d\n", *devPort)
		err = grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	// Human Bridge
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	bridge := ergo.NewHumanBridge(engine, e)

	log.Printf("[For Humans] listening on :%d\n", *humanPort)
	err = bridge.LaunchServerForHumans(fmt.Sprintf(":%d", *humanPort))
	if err != nil {
		panic(err)
	}

}
