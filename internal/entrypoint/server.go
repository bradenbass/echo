package entrypoint

import (
	"google.golang.org/grpc"

	"github.com/bradenbass/echo/internal/api"
	echopb "github.com/bradenbass/echo/proto"
)

func NewAPIServer(grpcServer *grpc.Server) {
	// Create a new echo server
	echoServer := api.NewEchoServer()

	// Register our implementation of the echo server onto the gRPC handler
	echopb.RegisterEchoerServer(grpcServer, echoServer)
}
