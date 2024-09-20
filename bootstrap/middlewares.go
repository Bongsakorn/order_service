package bootstrap

import (
	"encoding/json"
	"order_service/internal/types"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type LoggerTemplate struct {
	Logger              string `json:"logger"`
	ProjectName         string `json:"projectName"`
	AppName             string `json:"appName"`
	CompanyName         string `json:"company"`
	Env                 string `json:"env"`
	Thread              string `json:"thread"`
	UserId              string `json:"userId"`
	TraceId             string `json:"traceId"`
	SpanId              string `json:"spanId"`
	IpAddress           string `json:"ipAddress"`
	ClassName           string `json:"className"`
	LogType             string `json:"logType"`
	Uri                 string `json:"uri"`
	HttpStatus          string `json:"httpStatus"`
	OverallResponseTime int64  `json:"overallResponseTime"`
	Message             string `json:"message"`
	LogLevel            string `json:"logLevel"`
	Timestamp           string `json:"timestamp"`
	RequestBody         string `json:"requestBody"`
	requestParams       string `json:"requestParams"`
	RequestForm         string `json:"requestForm"`
	responseBody        string `json:"responseBody"`
	method              string `json:"method"`
}

func NewLoggerTemplate() *LoggerTemplate {
	return &LoggerTemplate{
		Logger:              "https://gitlab.com/scbtechx/pv-robinhood/robinhood-backend/ev-api",
		ProjectName:         "order_service",
		AppName:             "order_service",
		CompanyName:         "PPV",
		Thread:              "${pid}",
		UserId:              "-",
		TraceId:             "-",
		SpanId:              "-",
		IpAddress:           "${ip}",
		ClassName:           "-",
		LogType:             "ApiLog",
		Uri:                 "${path}",
		HttpStatus:          "${status}",
		OverallResponseTime: 0,
		Message:             "-",
		LogLevel:            "INFO",
		Timestamp:           "${time}",
		requestParams:       "${queryParams}",
		RequestBody:         "${body}",
		RequestForm:         "${form:}",
		responseBody:        "${resBody}",
		method:              "${method}",
	}
}

func SetAppMiddleware(app *fiber.App) {
	app.Use(cors.New())
	app.Use(helmet.New())
	loggerTemplate := NewLoggerTemplate()
	logFormat, _ := json.Marshal(loggerTemplate)
	app.Use(logger.New(logger.Config{
		Next: func(c *fiber.Ctx) bool {
			path := c.Path()

			if path == "/healthz" {
				return true
			}
			return false
		},
		Format:     string(logFormat) + "\n",
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))
	app.Use(responseMiddleware)
}

func responseMiddleware(ctx *fiber.Ctx) error {
	err := ctx.Next()
	ctx.Response().Header.Set("Content-Type", "application/json")
	if err != nil {
		ctx.Status(err.(*fiber.Error).Code).JSON(&types.BaseResponse{
			Code:   fiber.StatusOK,
			Status: "success",
			Data:   err.(*fiber.Error).Message,
		})
		return nil
	}

	response := ctx.Locals("response")

	if response != nil {
		ctx.Status(fiber.StatusOK).JSON(&types.BaseResponse{
			Code:   fiber.StatusOK,
			Status: "success",
			Data:   response,
		})
		return nil
	}

	return nil
}
