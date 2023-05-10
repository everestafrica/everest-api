package controller

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	util "github.com/everestafrica/everest-api/internal/commons/util"
	"github.com/everestafrica/everest-api/internal/handlers"
	"github.com/everestafrica/everest-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

type ISettingsController interface {
	CreateCustomCategory(ctx *fiber.Ctx) error
	DeleteCustomCategory(ctx *fiber.Ctx) error

	CreatePriceAlert(ctx *fiber.Ctx) error
	DeletePriceAlert(ctx *fiber.Ctx) error

	CreateNewsInterest(ctx *fiber.Ctx) error
	DeleteNewsInterest(ctx *fiber.Ctx) error
	RegisterRoutes(app *fiber.App)
}

type settingsController struct {
	settingsService service.ISettingsService
}

func NewSettingsController() ISettingsController {
	return &settingsController{
		settingsService: service.NewSettingsService(),
	}
}

func (ctl *settingsController) RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	settings := v1.Group("/settings")

	settings.Post("/custom-category", handlers.SecureAuth(), ctl.CreateCustomCategory)
	settings.Delete("/custom-category/:id", handlers.SecureAuth(), ctl.DeleteCustomCategory)

	settings.Post("/price-alert", handlers.SecureAuth(), ctl.CreatePriceAlert)
	settings.Delete("/price-alert/:id", handlers.SecureAuth(), ctl.DeletePriceAlert)

	settings.Post("/news-interest", handlers.SecureAuth(), ctl.CreateNewsInterest)
	settings.Delete("/news-interest/:id", handlers.SecureAuth(), ctl.DeleteNewsInterest)

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
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	categoryId := ctx.Params("id")
	err = ctl.settingsService.DeleteCustomCategory(userId, categoryId)
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

func (ctl *settingsController) CreatePriceAlert(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	var body *types.CreatePriceAlert
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

	err = ctl.settingsService.CreatePriceAlert(body, userId)
	if err != nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "price alert successfully added",
	})
}

func (ctl *settingsController) DeletePriceAlert(ctx *fiber.Ctx) error {
	_, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	alertId := ctx.Params("id")
	err = ctl.settingsService.DeletePriceAlert(alertId)
	if err != nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: "unable to delete price alert",
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "price alert successfully deleted",
	})
}

func (ctl *settingsController) CreateNewsInterest(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	var body *[]types.AddNewsInterest
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

	err = ctl.settingsService.CreateNewsInterest(body, userId)
	if err != nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "news interests (categories) successfully added",
	})
}

func (ctl *settingsController) DeleteNewsInterest(ctx *fiber.Ctx) error {
	_, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}

	interestId := ctx.Params("id")
	err = ctl.settingsService.DeletePriceAlert(interestId)
	if err != nil {
		return ctx.JSON(types.GenericResponse{
			Success: false,
			Message: "unable to delete interest",
		})
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "interest successfully deleted",
	})
}
