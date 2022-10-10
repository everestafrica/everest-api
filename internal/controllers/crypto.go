package controllers

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	util "github.com/everestafrica/everest-api/internal/commons/utils"
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
	crypto.Delete("/wallet", handlers.SecureAuth(), ctl.UnLinkWallet)
}

func (ctl *cryptoController) LinkWallet(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	var body types.CryptoWalletRequest

	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: "Problem while parsing request body",
		})
	}
	errors := util.ValidateStruct(body)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)

	}

	err = ctl.cryptoDetailsService.AddWallet(types.CryptoSymbol(body.Symbol), body.Address, userId)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "successfully linked user wallet",
	})
}

func (ctl *cryptoController) UnLinkWallet(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	var body types.CryptoWalletRequest

	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: "Problem while parsing request body",
		})
	}
	errors := util.ValidateStruct(body)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)

	}

	err = ctl.cryptoDetailsService.DeleteWallet(types.CryptoSymbol(body.Symbol), body.Address, userId)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "successfully unlinked user wallet",
	})
}
