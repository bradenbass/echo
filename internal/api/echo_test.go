package api

import (
	"context"
	"testing"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	echopb "github.com/bradenbass/echo/proto"
)

func TestEchoServer_Echo(t *testing.T) {
	// Arrange
	echoServer := &EchoServer{}
	expect := &echopb.EchoResponse{Reply: "Echo!!"}

	// Act
	res, err := echoServer.Echo(context.Background(), &echopb.EchoRequest{
		Message: "Echo!!",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	if !proto.Equal(expect, res) {
		t.Errorf("Expected to equal %v but got %v", expect, res)
	}
}

func TestEchoServer_Echo_RequiresMessage(t *testing.T) {
	// Arrange
	echoServer := &EchoServer{}
	expectedCode := codes.InvalidArgument
	expectedErrMessage := "No message to echo back"

	// Act
	res, err := echoServer.Echo(context.Background(), &echopb.EchoRequest{
		Message: "",
	})

	// Assert
	if res != nil {
		t.Fatalf("unexpected res received %v", err)
	}
	statusErr, ok := status.FromError(err)
	if !ok {
		t.Fatalf("unexpected error to be a Status but got %v", err)
	}
	if statusErr.Code() != expectedCode {
		t.Errorf("Expected to equal %v but got %v", expectedCode, statusErr.Code())
	}
	if statusErr.Message() != expectedErrMessage {
		t.Errorf("Expected to equal %v but got %v", expectedErrMessage, statusErr.Message())
	}
}
