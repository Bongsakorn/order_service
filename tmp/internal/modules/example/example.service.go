package example

import (
	"context"
	"log"
	"order_service/internal/config"
	pb "order_service/internal/grpc/protobuf"
	"order_service/internal/repository"
	"order_service/pkg/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ExampleServiceInterface interface {
	GetExample() (string, error)
	SayHello() (string, error)
	SayHi() (string, error)
}

type ExampleService struct {
	db          *repository.Repository
	client      *http.HqRentalClient
	grpcClient  pb.GreeterClient
	grpcClient2 pb.GreetingServiceClient
}

func NewService(repo *repository.Repository, hqClient *http.HqRentalClient, config *config.ConfigService) ExampleServiceInterface {
	println("service running...", config)
	conn, err := grpc.Dial(config.GreeterAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn2, err2 := grpc.Dial(config.GreetingAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil || err2 != nil {
		log.Fatalf("did not connect: %v", err)
		panic(err)
	}
	client := pb.NewGreeterClient(conn)
	client2 := pb.NewGreetingServiceClient(conn2)
	return &ExampleService{
		db:          repo,
		client:      hqClient,
		grpcClient:  client,
		grpcClient2: client2,
	}
}

func (service *ExampleService) GetExample() (string, error) {
	println("service running...")
	return "Hello, World!", nil
}

func (service *ExampleService) SayHello() (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := service.grpcClient.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	service.grpcClient.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	return r.GetMessage(), err
}

func (service *ExampleService) SayHi() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := service.grpcClient2.SayHi(ctx, &pb.HiRequest{Name: "world"})
	return r.GetMessage(), err
}
