package controllers

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/handlers"
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type IAccountController interface {
	RegisterRoutes(app *fiber.App)
	LinkAccount(ctx *fiber.Ctx) error
}

type accountController struct {
	accountService services.IAccountService
}

// NewAccountController instantiates Account Controller
func NewAccountController() IAccountController {
	return &accountController{
		accountService: services.NewAccountService(),
	}
}

func (ctl accountController) RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	accounts := v1.Group("/accounts")
	accounts.Post("/connect", handlers.SecureAuth(), ctl.LinkAccount)
	accounts.Post("/connect", handlers.SecureAuth(), ctl.LinkAccount)
}

func (ctl accountController) LinkAccount(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	var body *types.MonoAccountIdRequest
	//body := new(types.MonoAccountIdRequest)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: "Problem while parsing request body",
		})
	}
	err = ctl.accountService.SetAccountDetails(body.Code, userId)
	if err != nil {
		return err
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "successfully linked user account",
	})
}

func (ctl accountController) UnLinkAccount(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	accountId := ctx.Params("id")
	//if err := ctx.BodyParser(body); err != nil {
	//	return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
	//		Success: false,
	//		Message: "Problem while parsing request body",
	//	})
	//}
	err = ctl.accountService.UnlinkAccount(accountId, userId)
	if err != nil {
		return err
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "successfully unlinked user account",
	})
}
