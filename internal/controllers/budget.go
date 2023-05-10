package controllers

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	util "github.com/everestafrica/everest-api/internal/commons/utils"
	"github.com/everestafrica/everest-api/internal/handlers"
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type IBudgetController interface {
	GetBudget(ctx *fiber.Ctx) error
	AddBudget(ctx *fiber.Ctx) error
	UpdateBudget(ctx *fiber.Ctx) error
	DeleteBudget(ctx *fiber.Ctx) error
	RegisterRoutes(app *fiber.App)
}

type budgetController struct {
	budgetService services.IBudgetService
}

func NewBudgetController() IBudgetController {
	return &budgetController{
		budgetService: services.NewBudgetService(),
	}
}

func (ctl *budgetController) RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	budget := v1.Group("/budgets")

	budget.Get("/", handlers.SecureAuth(), ctl.GetBudget)
	budget.Post("/", handlers.SecureAuth(), ctl.AddBudget)
	budget.Put("/", handlers.SecureAuth(), ctl.UpdateBudget)
	budget.Delete("/", handlers.SecureAuth(), ctl.DeleteBudget)

}

func (ctl *budgetController) GetBudget(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	month := ctx.Query("month")
	stryear := ctx.Query("year")
	year, _ := strconv.Atoi(stryear)

	if len(month) < 1 {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: "required parameter `month` missing",
		})
	}

	if len(stryear) < 1 {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: "required parameter `year` missing",
		})
	}

	budget, err := ctl.budgetService.GetBudget(month, year, userId)
	if err != nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: "unable to get budget",
		})
	}
	if budget == nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: "unable to get budget, ensure query parameters are correct and try again",
		})
	}

	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "budget successfully retrieved",
		Data:    budget,
	})
}

func (ctl *budgetController) AddBudget(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	var body *types.CreateBudgetRequest
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

	err = ctl.budgetService.CreateBudget(body, userId)
	if err != nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "budget successfully added",
	})
}

func (ctl *budgetController) UpdateBudget(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	var body *types.UpdateBudgetRequest
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
	err = ctl.budgetService.UpdateBudget(body, userId)
	if err != nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: "unable to update budget",
		})
	}

	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "budget successfully updated",
	})
}

func (ctl *budgetController) DeleteBudget(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	month := ctx.Query("month")
	stryear := ctx.Query("year")
	year, _ := strconv.Atoi(stryear)

	if len(month) < 1 {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: "required parameter `month` missing",
		})
	}

	if len(stryear) < 1 {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: "required parameter `year` missing",
		})
	}

	err = ctl.budgetService.DeleteBudget(month, year, userId)
	if err != nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: "unable to delete budget",
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "budget successfully deleted",
	})
}
