package controllers

import (
	"github.com/everestafrica/everest-api/internal/commons/constants"
	"github.com/everestafrica/everest-api/internal/commons/types"
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type IAuthController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	RegisterRoutes(app *fiber.App)
}

type authController struct {
	authService services.IAuthService
}

// NewAuthController instantiates Auth Controller
func NewAuthController() IAuthController {
	return &authController{
		authService: services.NewAuthService(),
	}
}

func (ctl authController) RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	auth := v1.Group("/auth")
	auth.Post("/register", ctl.Register)
	auth.Post("/login", ctl.Login)
}

func (ctl authController) Register(ctx *fiber.Ctx) error {
	var body types.RegisterRequest
	requestIdentifier := uuid.NewString()

	ctx.Set(constants.RequestIdentifier, requestIdentifier)

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err,
		})
	}

	res, err := ctl.authService.Register(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err,
		})
	}

	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "User successfully registered",
		Data:    res,
	})
}
func (ctl authController) Login(ctx *fiber.Ctx) error {
	var body types.LoginRequest
	requestIdentifier := uuid.NewString()

	ctx.Set(constants.RequestIdentifier, requestIdentifier)

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err,
		})
	}

	res, err := ctl.authService.Login(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err,
		})
	}

	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "User successfully logged in",
		Data:    res,
	})
}
