package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigService struct {
	GoEnv                        string
	Port                         string
	DbUsername                   string
	DbPassword                   string
	DbHost                       string
	DbPort                       string
	DbName                       string
	HqUrl                        string
	HqTenantToken                string
	HqUserToken                  string
	GreetingAddress              string
	GreeterAddress               string
	GrpcPort                     int32
	AzServiceBusConnectionString string
}

func NewConfigService( /* for test propose we can inject vault service here */ ) *ConfigService {
	env := os.Getenv("GO_ENV")
	godotenv.Load(".env.local")
	if env != "local" {
		// Load from vault here
		godotenv.Load(fmt.Sprintf(".env.%s", env))
	}

	grpcPort, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		panic(err)
	}
	return &ConfigService{
		GoEnv:                        os.Getenv("GO_ENV"),
		Port:                         os.Getenv("PORT"),
		DbUsername:                   os.Getenv("DB_USERNAME"),
		DbPassword:                   os.Getenv("DB_PASSWORD"),
		DbHost:                       os.Getenv("DB_HOST"),
		DbPort:                       os.Getenv("DB_PORT"),
		DbName:                       os.Getenv("DB_NAME"),
		HqUrl:                        os.Getenv("HQ_URL"),
		HqTenantToken:                os.Getenv("HQ_TENANT_TOKEN"),
		HqUserToken:                  os.Getenv("HQ_USER_TOKEN"),
		GrpcPort:                     int32(grpcPort),
		AzServiceBusConnectionString: os.Getenv("SERVICEBUS_CONNECTION_STRING"),
	}
}
