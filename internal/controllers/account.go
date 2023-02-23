package controllers

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	util "github.com/everestafrica/everest-api/internal/commons/utils"
	"github.com/everestafrica/everest-api/internal/handlers"
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type IAccountController interface {
	RegisterRoutes(app *fiber.App)
	LinkAccount(ctx *fiber.Ctx) error
	UnLinkAccount(ctx *fiber.Ctx) error
	ReauthoriseUser(ctx *fiber.Ctx) error
}

type accountController struct {
	accountDetailsService services.IAccountDetailsService
}

// NewAccountController instantiates Account Controller
func NewAccountController() IAccountController {
	return &accountController{
		accountDetailsService: services.NewAccountDetailsService(),
	}
}

func (ctl *accountController) RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	accounts := v1.Group("/accounts")
	accounts.Post("/connect", ctl.LinkAccount)
	accounts.Post("/disconnect", handlers.SecureAuth(), ctl.UnLinkAccount)
	accounts.Get("/reauth", handlers.SecureAuth(), ctl.ReauthoriseUser)
}

func (ctl *accountController) LinkAccount(ctx *fiber.Ctx) error {
	//userId, err := handlers.UserFromContext(ctx)
	//if err != nil {
	//	return err
	//}
	userId := "1"

	var body types.MonoAccountIdRequest
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	errors := util.ValidateStruct(body)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)

	}

	err := ctl.accountDetailsService.SetAccountDetails(body.Code, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "successfully linked user account",
	})
}

func (ctl *accountController) ReauthoriseUser(ctx *fiber.Ctx) error {
	return nil
}

func (ctl *accountController) UnLinkAccount(ctx *fiber.Ctx) error {

	accountId := ctx.Params("id")
	err := ctl.accountDetailsService.UnlinkAccount(accountId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "successfully unlinked user account",
	})
}

func (ctl *accountController) GetAllTransactions(ctx *fiber.Ctx) error {

}

func (ctl *accountController) GetTransaction(ctx *fiber.Ctx) error {

}
