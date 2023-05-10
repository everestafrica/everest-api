package controllers

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	util "github.com/everestafrica/everest-api/internal/commons/utils"
	"github.com/everestafrica/everest-api/internal/handlers"
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type IAuthController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	SendOTP(ctx *fiber.Ctx) error
	//SendCode(ctx *fiber.Ctx) error
	RefreshToken(ctx *fiber.Ctx) error
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
	//auth.Post("/send-code", ctl.SendCode)
	auth.Post("/send-otp", ctl.SendOTP)
	auth.Post("/refresh-token", handlers.SecureAuth(), ctl.RefreshToken)
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
func (ctl *authController) Login(ctx *fiber.Ctx) error {
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

func (ctl *authController) ResetPassword(ctx fiber.Ctx) error {
	return nil
}

func (ctl *authController) SendOTP(ctx *fiber.Ctx) error {
	var body types.SendCodeRequest

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	code, err := ctl.authService.SendEmailOTP(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return ctx.JSON(types.GenericResponse{
		Success: true,
		Message: "OTP sent successfully",
		Data:    code,
	})
}

//func (ctl *authController) SendCode(ctx *fiber.Ctx) error {
//	var body types.SendCodeRequest
//
//	if err := ctx.BodyParser(&body); err != nil {
//		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
//			Success: false,
//			Message: err.Error(),
//		})
//	}
//
//	err := ctl.authService.SendEmailOTP(&body)
//	if err != nil {
//		return ctx.Status(fiber.StatusBadRequest).JSON(types.GenericResponse{
//			Success: false,
//			Message: err.Error(),
//		})
//	}
//
//	return ctx.JSON(types.GenericResponse{
//		Success: true,
//		Message: "OTP sent successfully",
//	})
//}

func (ctl authController) RefreshToken(ctx *fiber.Ctx) error {
	_, err := handlers.UserFromContext(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(types.GenericResponse{
			Success: false,
			Message: "Unauthorized User",
		})
	}
	return nil
}
