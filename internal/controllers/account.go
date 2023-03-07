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
	GetAllTransactions(ctx *fiber.Ctx) error
	GetTransaction(ctx *fiber.Ctx) error
	CreateManualTransaction(ctx *fiber.Ctx) error
	DeleteManualTransaction(ctx *fiber.Ctx) error
}

type accountController struct {
	accountDetailsService     services.IAccountDetailsService
	accountTransactionService services.IAccountTransactionService
}

// NewAccountController instantiates Account Controller
func NewAccountController() IAccountController {
	return &accountController{
		accountDetailsService:     services.NewAccountDetailsService(),
		accountTransactionService: services.NewAccountTransactionService(),
	}
}

func (ctl *accountController) RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	accounts := v1.Group("/accounts")
	accounts.Post("/connect", ctl.LinkAccount)
	accounts.Post("/disconnect", handlers.SecureAuth(), ctl.UnLinkAccount)
	accounts.Get("/reauth", handlers.SecureAuth(), ctl.ReauthoriseUser)

	// transactions
	accounts.Get("/transactions", handlers.SecureAuth(), ctl.GetAllTransactions)
	accounts.Get("/transactions/:id", handlers.SecureAuth(), ctl.GetTransaction)
	accounts.Post("/transactions", handlers.SecureAuth(), ctl.CreateManualTransaction)
	accounts.Delete("/transactions/:id", handlers.SecureAuth(), ctl.DeleteManualTransaction)
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
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	page := ctx.QueryInt("page")
	size := ctx.QueryInt("size")

	pagination := types.Pagination{
		Page: page,
		Size: size,
	}

	transactions, err := ctl.accountTransactionService.GetAllTransactions(userId, pagination)
	if err != nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "Transactions successfully fetched",
		Data:    transactions,
	})

}

func (ctl *accountController) GetTransaction(ctx *fiber.Ctx) error {
	_, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	transactionId := ctx.Params("id")

	transaction, err := ctl.accountTransactionService.GetTransaction(transactionId)
	if err != nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "transaction successfully fetched",
		Data:    transaction,
	})
}

func (ctl *accountController) CreateManualTransaction(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}
	var body *types.CreateTransactionRequest
	if err = ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: "Problem while parsing request body",
			Data:    err.Error(),
		})
	}
	errors := util.ValidateStruct(body)
	if errors != nil {
		return ctx.JSON(errors)
	}
	err = ctl.accountTransactionService.CreateManualTransaction(userId, body)
	if err != nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: "Problem while creating transaction",
			Data:    err.Error(),
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "successfully created manual transaction",
	})
}

func (ctl accountController) DeleteManualTransaction(ctx *fiber.Ctx) error {
	_, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}
	transactionId := ctx.Params("id")
	err = ctl.accountTransactionService.DeleteManualTransaction(transactionId)
	if err != nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "transaction successfully deleted",
	})
}
