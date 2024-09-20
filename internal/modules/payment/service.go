package payment

import (
	"encoding/base64"
	"fmt"
	"order_service/internal/config"
	"order_service/internal/modules/azservicebus"
	"strings"

	"github.com/google/uuid"
)

type PaymentServiceInterface interface {
	PaymentSuccess() error
	PaymentFailed() error
}

type PaymentService struct {
	azservicebus azservicebus.AzServiceBusServiceInterface
}

func NewService(config *config.ConfigService, servicebus azservicebus.AzServiceBusServiceInterface) PaymentServiceInterface {
	return &PaymentService{
		azservicebus: servicebus,
	}
}

// paymentFailed implements PaymentServiceInterface.
func (p *PaymentService) PaymentFailed() error {
	oid, err := GenerateShortUUID()
	if err != nil {
		return err
	}
	payload := fmt.Sprintf(`{"order_id": "%s","payment_status":"success","update_order_status":"failed"}`, oid)
	if err := p.azservicebus.SendMessage("test_q", payload); err != nil {
		return err

	}
	return nil
}

// paymentSuccess implements PaymentServiceInterface.
func (p *PaymentService) PaymentSuccess() error {
	oid, err := GenerateShortUUID()
	if err != nil {
		return err
	}
	payload := fmt.Sprintf(`{"order_id": "%s","payment_status":"success","update_order_status":"success"}`, oid)
	if err := p.azservicebus.SendMessage("test_q", payload); err != nil {
		return err
	}
	return nil
}

// GenerateShortUUID generates a short UUID
func GenerateShortUUID() (string, error) {
	// Generate a new UUID
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	// Encode the UUID using Base64 URL encoding
	shortUUID := base64.URLEncoding.EncodeToString(u[:])

	// Remove any padding characters ('=')
	shortUUID = strings.TrimRight(shortUUID, "=")

	return shortUUID, nil
}
