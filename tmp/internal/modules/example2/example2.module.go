package example2

import (
	"order_service/internal/config"
	"order_service/internal/modules/example"
	"order_service/internal/repository"
	"order_service/pkg/http"
)

type ExampleModule struct {
	ExampleService    ExampleService2Interface
	ExampleController ExampleController2Interface
}

func NewModule(repo *repository.Repository, config *config.ConfigService) *ExampleModule {
	hqClient := http.NewHqRentalClient("", "", "")
	service := NewService(example.NewService(repo, hqClient, config))
	module := &ExampleModule{
		ExampleService:    service,
		ExampleController: NewController(service),
	}
	return module

}
