package controllers

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/handlers"
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/gofiber/fiber/v2"
	"log"
)

type IWebhookController interface {
	Mono(ctx *fiber.Ctx) error
	RegisterRoutes(app *fiber.App)
}

type webhookController struct {
	webhookService services.IWebhookService
}

func NewWebhookController() IWebhookController {
	return &webhookController{
		webhookService: services.NewWebhookService(),
	}
}

func (ctl webhookController) RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/v1/webhook")
	v1.Post("/mono", handlers.VerifyMonoWebhook(), ctl.Mono)

}

func (ctl webhookController) Mono(ctx *fiber.Ctx) error {

	var body types.MonoWebhookPayload
	if err := ctx.BodyParser(body); err != nil {
		log.Print(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: "Problem while parsing webhook body",
		})
	}
	err := ctl.webhookService.MonoWebhook(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(types.GenericResponse{
		Success: true,
		Message: "webhook received successfully",
		Data:    body,
	})
}
