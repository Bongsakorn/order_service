package grpcexample

import (
	pb "order_service/internal/grpc/protobuf"
	"testing"
)

func TestGetGrpc(t *testing.T) {

	server := Server{}

	server.SayHello(nil, &pb.HelloRequest{
		Name: "test",
	})

}
