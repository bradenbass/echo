package e2e

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"

	echo "github.com/bradenbass/echo/client/sdk"
	"github.com/bradenbass/echo/internal/entrypoint"
	echopb "github.com/bradenbass/echo/proto"
)

func Test_e2e(t *testing.T) {
	// Arrange
	expectedResponse := &echopb.EchoResponse{Reply: "Hello!"}

	// Listen on a random tcp port
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Error listening: %v", err)
	}
	// Start a new gRPC Server
	grpcServer := grpc.NewServer()
	entrypoint.NewAPIServer(grpcServer)

	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			t.Fatal(err)
		}
	}()
	defer grpcServer.Stop()

	client, err := echo.NewClient(lis.Addr().String(), false)
	if err != nil {
		t.Fatalf("Unable to create a new client %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Act
	res, err := client.Echo(ctx, &echopb.EchoRequest{Message: "Hello!"}, grpc.WaitForReady(true))

	// Assert
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !proto.Equal(expectedResponse, res) {
		t.Errorf("Expected response %v but got back %v", expectedResponse, res)
	}
}
