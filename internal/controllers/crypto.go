package controllers

import (
	"github.com/everestafrica/everest-api/internal/handlers"
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ICryptoController interface {
	RegisterRoutes(app *fiber.App)
	LinkWallet(ctx *fiber.Ctx) error
}

type cryptoController struct {
	cryptoDetailsService services.ICryptoService
}

// NewCryptoController instantiates Crypto Controller
func NewCryptoController() ICryptoController {
	return &cryptoController{
		cryptoDetailsService: services.NewCryptoService(),
	}
}

func (ctl *cryptoController) RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	crypto := v1.Group("/crypto")
	crypto.Post("/wallet", handlers.SecureAuth(), ctl.LinkWallet)
}

func (ctl *cryptoController) LinkWallet(ctx *fiber.Ctx) error {
	return nil
}
