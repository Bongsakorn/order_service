package grpcexample

import (
	"context"
	"log"

	pb "order_service/internal/grpc/protobuf"
	"order_service/internal/repository"
	repositoryModel "order_service/internal/repository/models"
)

type Server struct {
	pb.UnimplementedGreeterServer
	Repository *repository.Repository
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	var fleetsResponse []repositoryModel.Fleet
	s.Repository.Db.Model(&repositoryModel.Fleet{}).Where("active = ?", true).Find(&fleetsResponse)

	println(fleetsResponse)
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}
