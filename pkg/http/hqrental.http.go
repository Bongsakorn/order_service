package http

import (
	"encoding/base64"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type HqRentalClient struct {
	HttpClient *resty.Client
}

func NewHqRentalClient(baseUrl string, tenantToken string, userToken string) *HqRentalClient {
	token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", tenantToken, userToken)))
	return &HqRentalClient{
		HttpClient: resty.New().SetBaseURL(baseUrl).SetHeader("Authorization", fmt.Sprintf("Basic %s", token)),
	}
}
