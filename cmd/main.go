package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/bregydoc/ergo"
	"github.com/bregydoc/ergo/client"
	"github.com/bregydoc/ergo/creators"
	ergocon "github.com/bregydoc/ergo/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
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

	// Ergo Engine
	server := ergo.NewErgoServer(e)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *devPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	ergocon.RegisterErgoServer(grpcServer, server)

	go func() {
		log.Printf("[For Developers] GRPC listening on :%d\n", *devPort)
		err = grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	// Human Bridge
	gin.SetMode(gin.ReleaseMode)
	humanEngine := gin.Default()
	bridge := ergo.NewHumanBridge(humanEngine, e)

	go func() {
		log.Printf("[For Humans] Human Wizard listening on :%d\n", *humanPort)
		err = bridge.LaunchServerForHumans(fmt.Sprintf(":%d", *humanPort))
		if err != nil {
			panic(err)
		}
	}()

	// UI Bridge
	uiEngine := gin.Default()
	ui := client.NewErgoUI(uiEngine, e)
	log.Printf("[For Developers] UI listening on :5000\n")
	err = ui.LaunchUIClientForDevelopers(":5000")
	if err != nil {
		panic(err)
	}
}
