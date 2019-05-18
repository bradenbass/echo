package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"

	echopb "github.com/bradenbass/echo/proto"
)

func main() {
	args := os.Args

	// Create a new Echoer Client
	echoerClient, err := createClient()
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

func createClient() (echopb.EchoerClient, error) {
	// Dial a new insecure connection
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return echopb.NewEchoerClient(conn), nil
}
