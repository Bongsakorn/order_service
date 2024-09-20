package example2

import (
	"order_service/internal/types"

	"github.com/gofiber/fiber/v2"
)

type ExampleController2Interface interface {
	fetchExample(ctx *fiber.Ctx) error
	GetRoute() types.Route
}

type ExampleController2 struct {
	ExampleService ExampleService2Interface
}

func NewController(service ExampleService2Interface) ExampleController2Interface {
	return &ExampleController2{ExampleService: service}
}

func (controller *ExampleController2) GetRoute() types.Route {
	return types.Route{
		Version: "1",
		Path:    "/example",
		Children: []types.Children{
			{
				Method:  "GET",
				Path:    "/extra",
				Handler: controller.fetchExample,
			},
		},
	}
}
func (controller *ExampleController2) fetchExample(ctx *fiber.Ctx) error {
	res, err := controller.ExampleService.GetExample()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": res})
}
