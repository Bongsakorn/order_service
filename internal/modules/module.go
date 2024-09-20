package module

import (
	"order_service/internal/config"
	"order_service/internal/types"
)

// NewModule creates a new module with the given config and returns a list of routes.
func NewModule(config *config.ConfigService) []types.Route {
	// repository := repository.NewRepository(config)
	// exampleModule := example.NewModule(config, repository)
	// example2Module := example2.NewModule(repository, config)
	// azServiceBusModule := azservicebus.NewModule(config)
	// paymentModule := payment.NewModule(config, repository, azServiceBusModule)

	routes := []types.Route{}

	return routes
}
