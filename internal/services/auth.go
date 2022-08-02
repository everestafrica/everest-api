package services

import (
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/everestafrica/everest-api/internal/repositories"
)

type IAuthService interface {
}

type authService struct {
	jwtSecret string
	userRepo  repositories.IUserRepository
}

// NewAuthService will instantiate AuthService
func NewAuthService() IAuthService {
	return &authService{
		jwtSecret: config.GetConf().JWTSecret,
		userRepo:  repositories.NewUserRepo(),
	}
}
