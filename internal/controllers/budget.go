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

func (bc *budgetController) RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/v1/budget")
	
	v1.Get("/", handlers.SecureAuth(), bc.GetBudget)
	v1.Post("/", handlers.SecureAuth(), bc.AddBudget)
	v1.Put("/", handlers.SecureAuth(), bc.UpdateBudget)
	v1.Delete("/", handlers.SecureAuth(), bc.DeleteBudget)

}

func (bc *budgetController) GetBudget(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	month := ctx.Query("month")
	year, _ := strconv.Atoi(ctx.Query("year"))

	budget, err := bc.budgetService.GetBudget(month, year, userId)
	if err != nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: "unable to get budget",
		})
	}

	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "budget successfully added",
		Data:    budget,
	})
}

func (bc *budgetController) AddBudget(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	var body *types.CreateBudgetRequest
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: "Problem while parsing request body",
		})
	}

	errors := util.ValidateStruct(body)
	if errors != nil {
		return ctx.JSON(errors)
	}

	err = bc.budgetService.CreateBudget(body, userId)
	if err != nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: "unable to add budget",
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "budget successfully added",
		Data:    body,
	})
}

func (bc *budgetController) UpdateBudget(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	var body *types.UpdateBudgetRequest
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: "Problem while parsing request body",
		})
	}

	errors := util.ValidateStruct(body)
	if errors != nil {
		return ctx.JSON(errors)
	}
	err = bc.budgetService.UpdateBudget(body, userId)
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

func (bc *budgetController) DeleteBudget(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	month := ctx.Query("month")
	year, _ := strconv.Atoi(ctx.Query("year"))

	err = bc.budgetService.DeleteBudget(month, year, userId)
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
