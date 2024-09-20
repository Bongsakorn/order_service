package payment

import (
	"order_service/internal/config"
	"order_service/internal/modules/azservicebus"
	"order_service/internal/repository"
)

type PaymentModule struct {
	PaymentService    PaymentServiceInterface
	PaymentController PaymentControllerInterface
}

func NewModule(config *config.ConfigService, repo *repository.Repository, servicebus *azservicebus.AzServiceBusModule) *PaymentModule {
	service := NewService(config, servicebus.AzServiceBusService)
	module := &PaymentModule{
		PaymentService:    service,
		PaymentController: NewController(service),
	}
	return module
}
