package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"

	"github.com/bradenbass/echo/client/sdk"
	echopb "github.com/bradenbass/echo/proto"
)

func main() {
	args := os.Args

	// Create a new Echoer Client
	echoerClient, err := echo.NewClient("127.0.0.1:9000", true)
	if err != nil {
		log.Fatalf("Unable to create a new client: %v", err)
	}

	// Limit how long we wait for a reply from the server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Send RPC to Server
	res, err := echoerClient.Echo(ctx, &echopb.EchoRequest{
		Message: strings.Join(args[1:], " "),
	}, grpc.WaitForReady(true))

	if err != nil {
		log.Fatalf("Error trying to send message to server: %v", err)
	}
	log.Printf("Received reply from server: %s", res.Reply)
}
