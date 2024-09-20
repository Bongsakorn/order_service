package example_test

import (
	"order_service/internal/config"
	"order_service/internal/modules/example"
	"order_service/internal/repository"
	"order_service/pkg/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetExample(t *testing.T) {
	defer httpmock.DeactivateAndReset()
	client := http.NewHqRentalClient("test", "test", "test")
	httpmock.ActivateNonDefault(client.HttpClient.GetClient())

	repository, _ := repository.NewMockRepository()

	config := config.NewConfigService()

	example.NewService(repository, client, config)

}
