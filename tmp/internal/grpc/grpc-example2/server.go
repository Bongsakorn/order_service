package grpcexample2

import (
	"context"
	"log"

	pb "order_service/internal/grpc/protobuf"
)

type Server struct {
	pb.UnimplementedGreetingServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHi(ctx context.Context, in *pb.HiRequest) (*pb.HiReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HiReply{Message: "Hi" + in.GetName()}, nil
}
