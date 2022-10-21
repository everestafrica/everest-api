package controllers

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	util "github.com/everestafrica/everest-api/internal/commons/utils"
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type IAuthController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	SendOTPCode(ctx *fiber.Ctx) error
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

func (ctl *authController) RegisterRoutes(app *fiber.App) {
	auth := app.Group("/v1/auth")

	auth.Post("/register", ctl.Register)
	auth.Post("/login", ctl.Login)
	auth.Post("/send-otp", ctl.SendOTPCode)
}

func (ctl *authController) Register(ctx *fiber.Ctx) error {
	var body types.RegisterRequest

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	errors := util.ValidateStruct(body)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)

	}

	res, err := ctl.authService.Register(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
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

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	res, err := ctl.authService.Login(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "User successfully logged in",
		Data:    res,
	})
}

func (ctl authController) ResetPassword(ctx fiber.Ctx) error {
	return nil
}

func (ctl authController) SendOTPCode(ctx *fiber.Ctx) error {
	var body types.SendCodeRequest

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	err := ctl.authService.SendOTPCode(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "OTP sent successfully",
	})
}
