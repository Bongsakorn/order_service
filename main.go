package main

import (
	"order_service/bootstrap"

	"order_service/internal/config"
	"order_service/internal/grpc"

	module "order_service/internal/modules"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config := config.NewConfigService()
	routes := module.NewModule(config)
	go grpc.StartGrpcServer(config)
	bootstrap.Setup(app, config, routes)
}
