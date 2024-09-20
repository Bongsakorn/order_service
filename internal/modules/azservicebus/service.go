package azservicebus

import (
	"context"
	"fmt"
	"log"
	"order_service/internal/config"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

type AzServiceBusServiceInterface interface {
	GetClient() *azservicebus.Client
	SendMessage(queueName string, message string) error
}

type AzServiceBusService struct {
	client *azservicebus.Client
}

func NewService(config *config.ConfigService) AzServiceBusServiceInterface {
	conn, connErr := azservicebus.NewClientFromConnectionString(config.AzServiceBusConnectionString, nil)

	if connErr != nil {
		log.Fatalf("failed to create service bus client: %v", connErr)
		panic(connErr)
	}

	return &AzServiceBusService{
		client: conn,
	}
}

func (az *AzServiceBusService) GetClient() *azservicebus.Client {
	return az.client
}

// sendMessage implements AzServiceBusServiceInterface.
func (az *AzServiceBusService) SendMessage(queueName string, message string) error {
	// Create a sender for the queue
	sender, err := az.client.NewSender(queueName, nil)
	if err != nil {
		return fmt.Errorf("failed to create queue sender: %w", err)
	}

	payload := &azservicebus.Message{
		Body: []byte(message),
	}

	// Send the message
	ctx := context.Background()
	err = sender.SendMessage(ctx, payload, nil)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	fmt.Printf("message sent successfully to queue %s\n", queueName)
	return nil
}
