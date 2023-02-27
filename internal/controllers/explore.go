package controllers

import (
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type IExploreController interface {
	RegisterRoutes(app *fiber.App)
}

type exploreController struct {
	newsService services.INewsService
}

func (ctl *exploreController) GetUserNews(ctx *fiber.Ctx) error {
	return nil
}
