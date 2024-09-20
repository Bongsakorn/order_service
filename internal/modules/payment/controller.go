package payment

import (
	"fmt"
	"order_service/internal/types"

	"github.com/gofiber/fiber/v2"
)

type PaymentControllerInterface interface {
	GetRoute() types.Route
	paymentSuccess(ctx *fiber.Ctx) error
	paymentFailed(ctx *fiber.Ctx) error
}

type PaymentController struct {
	PaymentService PaymentServiceInterface
}

func NewController(service PaymentServiceInterface) PaymentControllerInterface {
	return &PaymentController{PaymentService: service}
}

func (controller *PaymentController) GetRoute() types.Route {
	return types.Route{
		Version: "1",
		Path:    "/payment",
		Children: []types.Children{
			{
				Method:  "GET",
				Path:    "/success",
				Handler: controller.paymentSuccess,
			},
			{
				Method:  "GET",
				Path:    "/failed",
				Handler: controller.paymentFailed,
			},
		},
	}
}

func (controller *PaymentController) paymentSuccess(ctx *fiber.Ctx) error {
	if err := controller.PaymentService.PaymentSuccess(); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("something went wrong: %v", err),
		}
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Payment success"})
	return nil
}
func (controller *PaymentController) paymentFailed(ctx *fiber.Ctx) error {
	if err := controller.PaymentService.PaymentFailed(); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("something went wrong: %v", err),
		}
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Payment failed"})
	return nil
}
