package example

import (
	"order_service/internal/config"
	"order_service/internal/repository"
	"order_service/pkg/http"
)

type ExampleModule struct {
	ExampleService    ExampleServiceInterface
	ExampleController ExampleControllerInterface
}

func NewModule(config *config.ConfigService, repo *repository.Repository) *ExampleModule {
	hqClient := http.NewHqRentalClient(config.HqUrl, config.HqTenantToken, config.HqUserToken)
	service := NewService(repo, hqClient, config)
	module := &ExampleModule{
		ExampleService:    service,
		ExampleController: NewController(service),
	}
	return module
}
