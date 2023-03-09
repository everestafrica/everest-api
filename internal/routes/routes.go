package routes

import (
	"github.com/everestafrica/everest-api/internal/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router *fiber.App) {
	controllers.NewAuthController().RegisterRoutes(router)
	controllers.NewSubscriptionController().RegisterRoutes(router)
	controllers.NewAccountController().RegisterRoutes(router)
	controllers.NewBudgetController().RegisterRoutes(router)
	controllers.NewWebhookController().RegisterRoutes(router)
	controllers.NewCryptoController().RegisterRoutes(router)
	controllers.NewDebtController().RegisterRoutes(router)
	controllers.NewAssetController().RegisterRoutes(router)
	controllers.NewSettingsController().RegisterRoutes(router)
}
