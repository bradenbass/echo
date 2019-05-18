package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"

	"github.com/bradenbass/echo/internal/api"
	echopb "github.com/bradenbass/echo/proto"
)

const (
	grpcPort = 9000
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new echo server
	echoServer := api.NewEchoServer()

	// Create a new gRPC Server
	grpcServer := grpc.NewServer()

	// Register our implementation of the echo server onto the gRPC handler
	echopb.RegisterEchoerServer(grpcServer, echoServer)

	// Start Go Server
	go func() {
		log.Printf("Listening for gRPC requests...")
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("Error listening")
		}
	}()

	// Block until we receive a signal to shutdown
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	log.Printf("Shutting down server...")

	// Tell gRPC Server to stop processing requests and block till pending and in flight are finished
	grpcServer.GracefulStop()
}
