package api

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/bradenbass/echo/proto"
)

func NewEchoServer() *EchoServer {
	return &EchoServer{}
}

type EchoServer struct {
}

func (e *EchoServer) Echo(ctx context.Context, req *echopb.EchoRequest) (*echopb.EchoResponse, error) {
	if req.GetMessage() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "No message to echo back")
	}
	log.Printf("Received messsage: %s", req.GetMessage())
	return &echopb.EchoResponse{
		Reply: req.GetMessage(),
	}, nil
}
