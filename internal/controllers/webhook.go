package controllers

import (
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/gofiber/fiber/v2"
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
	v1.Post("/mono")
}

func (ctl webhookController) Mono(ctx *fiber.Ctx) error {
	return nil
}
