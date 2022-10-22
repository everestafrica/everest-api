package controllers

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	util "github.com/everestafrica/everest-api/internal/commons/utils"
	"github.com/everestafrica/everest-api/internal/handlers"
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ISubscriptionController interface {
	RegisterRoutes(app *fiber.App)
	GetSubscription(ctx *fiber.Ctx) error
	GetAllSubscriptions(ctx *fiber.Ctx) error
	AddSubscription(ctx *fiber.Ctx) error
	DeleteSubscription(ctx *fiber.Ctx) error
}

type subscriptionController struct {
	subscriptionService services.ISubscriptionService
}

// NewSubscriptionController instantiates Subscription Controller
func NewSubscriptionController() ISubscriptionController {
	return &subscriptionController{
		subscriptionService: services.NewSubscriptionService(),
	}
}

func (ctl subscriptionController) RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	subs := v1.Group("/subscriptions")
	subs.Get("/:id", handlers.SecureAuth(), ctl.GetSubscription)
	subs.Get("/", handlers.SecureAuth(), ctl.GetAllSubscriptions)
	subs.Post("/", handlers.SecureAuth(), ctl.AddSubscription)
	subs.Delete("/:id", handlers.SecureAuth(), ctl.DeleteSubscription)
}

func (ctl subscriptionController) GetSubscription(ctx *fiber.Ctx) error {
	return nil
}

func (ctl subscriptionController) GetAllSubscriptions(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return err
	}
	subscriptions, err := ctl.subscriptionService.GetAllSubscriptions(userId)
	if err != nil {
		return err
	}

	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "Subscriptions retrieved",
		Data:    subscriptions,
	})
}
func (ctl subscriptionController) AddSubscription(ctx *fiber.Ctx) error {
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
			Success: false,
			Message: "Unauthorized User",
		})
	}

	var body types.SubscriptionRequest
	//body := new(types.SubscriptionRequest)
	if err = ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	errors := util.ValidateStruct(body)
	if errors != nil {
		return ctx.JSON(errors)
	}
	err = ctl.subscriptionService.AddSubscription(&body, userId)
	if err != nil {
		return err
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "Successfully created subscription",
	})
}

func (ctl subscriptionController) DeleteSubscription(ctx *fiber.Ctx) error {
	subId := ctx.Params("id")
	userId, err := handlers.UserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
			Success: false,
			Message: "Unauthorized User",
		})
	}
	err = ctl.subscriptionService.DeleteSubscription(subId, userId)
	if err != nil {
		return err
	}
	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "Successfully deleted subscription",
	})
}
