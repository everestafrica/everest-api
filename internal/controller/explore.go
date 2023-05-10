package controller

import (
	"github.com/everestafrica/everest-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

type IExploreController interface {
	RegisterRoutes(app *fiber.App)
}

type exploreController struct {
	newsService service.INewsService
}

func (ctl *exploreController) GetUserNews(ctx *fiber.Ctx) error {
	return nil
}
