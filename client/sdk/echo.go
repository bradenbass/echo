package echo

import (
	"google.golang.org/grpc"

	echopb "github.com/bradenbass/echo/proto"
)

func NewClient(address string) (echopb.EchoerClient, error) {
	// Dial a new insecure connection
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return echopb.NewEchoerClient(conn), nil
}
