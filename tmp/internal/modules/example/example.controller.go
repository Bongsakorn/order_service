package example

import (
	"encoding/json"
	"order_service/internal/types"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/validator.v2"
)

type ExampleControllerInterface interface {
	fetchExample(ctx *fiber.Ctx) error
	PostExample(ctx *fiber.Ctx) error
	GetRoute() types.Route
	CallSayHello(ctx *fiber.Ctx) error
	CallSayHi(ctx *fiber.Ctx) error
}

type ExampleController struct {
	ExampleService ExampleServiceInterface
}

func NewController(service ExampleServiceInterface) ExampleControllerInterface {
	return &ExampleController{ExampleService: service}
}

func (controller *ExampleController) GetRoute() types.Route {
	return types.Route{
		Version: "1",
		Path:    "/example",
		Children: []types.Children{
			{
				Method:  "GET",
				Path:    "/",
				Handler: controller.fetchExample,
			},
			{
				Method:  "POST",
				Path:    "/",
				Handler: controller.PostExample,
			},
			{
				Method:  "GET",
				Path:    "/hello",
				Handler: controller.CallSayHello,
			},
			{
				Method:  "GET",
				Path:    "/hi",
				Handler: controller.CallSayHi,
			},
		},
	}
}

func (controller *ExampleController) fetchExample(ctx *fiber.Ctx) error {
	res, err := controller.ExampleService.GetExample()
	if err != nil {
		return err
	}
	ctx.Locals("response", res)
	// ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": res})
	return nil
}

func (controller *ExampleController) PostExample(ctx *fiber.Ctx) error {
	rawBody := ctx.Body()
	var body types.ExampleRequest
	err := json.Unmarshal(rawBody, &body)
	if err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
		}
	}
	if errs := validator.Validate(body); errs != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
		}
	}
	println(body.Bar)
	return nil
}

func (controller *ExampleController) CallSayHello(ctx *fiber.Ctx) error {
	res, error := controller.ExampleService.SayHello()
	if error != nil {
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: error.Error(),
		}

	}
	ctx.Locals("response", res)
	return nil
}

func (controller *ExampleController) CallSayHi(ctx *fiber.Ctx) error {
	res, error := controller.ExampleService.SayHi()
	if error != nil {
		return error

	}
	ctx.Locals("response", res)
	return nil
}
