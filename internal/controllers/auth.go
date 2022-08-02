package controllers

import "github.com/everestafrica/everest-api/internal/services"

type IAuthController interface {
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
