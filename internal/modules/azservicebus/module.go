package azservicebus

import "order_service/internal/config"

type AzServiceBusModule struct {
	AzServiceBusService AzServiceBusServiceInterface
}

func NewModule(config *config.ConfigService) *AzServiceBusModule {
	service := NewService(config)
	module := &AzServiceBusModule{
		AzServiceBusService: service,
	}
	return module
}
