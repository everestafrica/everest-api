package controllers

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	util "github.com/everestafrica/everest-api/internal/commons/utils"
	"github.com/everestafrica/everest-api/internal/handlers"
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ISettingsController interface {
	CreateCustomCategory(ctx *fiber.Ctx) error
	DeleteCustomCategory(ctx *fiber.Ctx) error
	RegisterRoutes(app *fiber.App)
}

type settingsController struct {
	settingsService services.ISettingsService
}

func NewSettingsController() ISettingsController {
	return &settingsController{
		settingsService: services.NewSettingsService(),
	}
}

func (ctl *settingsController) RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	settings := v1.Group("/settings")

	settings.Post("/custom-category", handlers.SecureAuth(), ctl.CreateCustomCategory)
	settings.Delete("/custom-category/:id", handlers.SecureAuth(), ctl.DeleteCustomCategory)

}

func (ctl *settingsController) CreateCustomCategory(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	var body *types.CreateCustomCategory
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

	err = ctl.settingsService.CreateCustomCategory(body, userId)
	if err != nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "custom category successfully added",
	})
}

func (ctl *settingsController) DeleteCustomCategory(ctx *fiber.Ctx) error {
	_, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	categoryId := ctx.Params("id")
	err = ctl.settingsService.DeleteCustomCategory(categoryId)
	if err != nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: "unable to delete custom category",
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "custom category successfully deleted",
	})
}
