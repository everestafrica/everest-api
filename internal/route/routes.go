package route

import (
	"github.com/everestafrica/everest-api/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router *fiber.App) {
	controller.NewAuthController().RegisterRoutes(router)
	controller.NewSubscriptionController().RegisterRoutes(router)
	controller.NewAccountController().RegisterRoutes(router)
	controller.NewBudgetController().RegisterRoutes(router)
	controller.NewWebhookController().RegisterRoutes(router)
	controller.NewCryptoController().RegisterRoutes(router)
	controller.NewDebtController().RegisterRoutes(router)
	controller.NewAssetController().RegisterRoutes(router)
	controller.NewSettingsController().RegisterRoutes(router)
}
