package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/bradenbass/echo/internal/entrypoint"
)

const (
	grpcPort = 9000

	certFile = "tls/Client.crt"
	keyFile  = "tls/Client.key"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	tlsCreds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("Error loading tls config %v", err)
	}

	// Create a new gRPC Server with a maximum number of concurrent streams and TLS config
	grpcServer := grpc.NewServer(
		grpc.MaxConcurrentStreams(10000),
		grpc.Creds(tlsCreds),
	)

	// Apply configuration onto gRPC server
	entrypoint.NewAPIServer(grpcServer)

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
