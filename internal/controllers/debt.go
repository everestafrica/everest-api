package controllers

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	util "github.com/everestafrica/everest-api/internal/commons/utils"
	"github.com/everestafrica/everest-api/internal/handlers"
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type IDebtController interface {
	RegisterRoutes(app *fiber.App)
	GetDebt(ctx *fiber.Ctx) error
	GetAllDebtsByType(ctx *fiber.Ctx) error
	AddDebt(ctx *fiber.Ctx) error
	DeleteDebt(ctx *fiber.Ctx) error
}

type debtController struct {
	debtService services.IDebtService
}

// NewDebtController instantiates Debt Controller
func NewDebtController() IDebtController {
	return &debtController{
		debtService: services.NewDebtService(),
	}
}

func (ctl debtController) RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	debt := v1.Group("/debts")
	debt.Get("/:id", handlers.SecureAuth(), ctl.GetDebt)
	debt.Get("/", handlers.SecureAuth(), ctl.GetAllDebtsByType)
	debt.Post("/", handlers.SecureAuth(), ctl.AddDebt)
	debt.Delete("/:id", handlers.SecureAuth(), ctl.DeleteDebt)
}

func (ctl debtController) GetDebt(ctx *fiber.Ctx) error {
	debtId, _ := strconv.Atoi(ctx.Params("id"))
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
			Success: false,
			Message: "Unauthorized User",
		})
	}
	debt, err := ctl.debtService.GetDebt(userId, debtId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: "Problem while retrieving debt",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(types.GenericResponse{
		Success: true,
		Message: "Debt retrieved",
		Data:    debt,
	})
}

func (ctl debtController) GetAllDebtsByType(ctx *fiber.Ctx) error {
	debtType := ctx.Query("type")
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
			Success: false,
			Message: "Unauthorized User",
		})
	}
	debts, err := ctl.debtService.GetDebtsByType(userId, types.DebtType(debtType))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: "Problem while retrieving debt",
		})
	}

	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "Debts retrieved",
		Data:    debts,
	})
}

func (ctl debtController) AddDebt(ctx *fiber.Ctx) error {
	var body *types.CreateDebtRequest
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
			Success: false,
			Message: "Unauthorized User",
		})
	}
	if err = ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: "Problem while parsing request body",
		})
	}
	errors := util.ValidateStruct(body)
	if errors != nil {
		return ctx.JSON(errors)
	}

	err = ctl.debtService.AddDebt(body, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: "Problem while creating debt",
		})
	}

	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "Debt created",
	})
}

func (ctl debtController) UpdateDebt(ctx *fiber.Ctx) error {
	var body *types.UpdateDebtRequest
	debtId, _ := strconv.Atoi(ctx.Params("id"))
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
			Success: false,
			Message: "Unauthorized User",
		})
	}
	err = ctl.debtService.UpdateDebt(body, userId, debtId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: "Problem while updating debt",
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "Debt updated",
	})
}

func (ctl debtController) DeleteDebt(ctx *fiber.Ctx) error {
	debtId := ctx.Params("id")
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
			Success: false,
			Message: "Unauthorized User",
		})
	}
	err = ctl.debtService.DeleteDebt(debtId, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: "Problem while deleting debt",
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "Debt deleted",
	})
}
