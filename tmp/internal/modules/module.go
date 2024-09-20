package module

import (
	"order_service/internal/config"
	"order_service/internal/modules/example"
	"order_service/internal/modules/example2"
	"order_service/internal/repository"
	"order_service/internal/types"
)

func NewModule(config *config.ConfigService) []types.Route {
	repository := repository.NewRepository(config)
	exampleModule := example.NewModule(config, repository)
	example2Module := example2.NewModule(repository, config)

	routes := []types.Route{exampleModule.ExampleController.GetRoute(), example2Module.ExampleController.GetRoute()}

	return routes

}
