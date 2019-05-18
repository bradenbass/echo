package api

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/bradenbass/echo/proto"
)

type EchoServer struct {
}

func (e *EchoServer) Echo(ctx context.Context, req *echopb.EchoRequest) (*echopb.EchoResponse, error) {
	if req.GetMessage() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "No message to echo back")
	}
	return &echopb.EchoResponse{
		Reply: req.GetMessage(),
	}, nil
}
