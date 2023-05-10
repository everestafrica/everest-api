package controllers

import (
	"github.com/everestafrica/everest-api/internal/handlers"
	"github.com/everestafrica/everest-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

type IAssetController interface {
	RegisterRoutes(app *fiber.App)
	AddAsset(ctx *fiber.Ctx) error
	GetAsset(ctx *fiber.Ctx) error
	GetAllAssets(ctx *fiber.Ctx) error
	DeleteAsset(ctx *fiber.Ctx) error
}

type assetController struct {
	assetService services.IAssetService
}

// NewAssetController instantiates assetController
func NewAssetController() IAssetController {
	return &assetController{
		assetService: services.NewAssetService(),
	}
}

func (ctl assetController) RegisterRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	asset := v1.Group("/assets")
	asset.Get("/:id", handlers.SecureAuth(), ctl.GetAsset)
	asset.Get("/", handlers.SecureAuth(), ctl.GetAllAssets)
	asset.Post("/", handlers.SecureAuth(), ctl.AddAsset)
	asset.Delete("/:id", handlers.SecureAuth(), ctl.DeleteAsset)
}

func (ctl assetController) AddAsset(ctx *fiber.Ctx) error {
	return nil
}

func (ctl assetController) GetAsset(ctx *fiber.Ctx) error {
	return nil
}
func (ctl assetController) GetAllAssets(ctx *fiber.Ctx) error {
	return nil
}
func (ctl assetController) DeleteAsset(ctx *fiber.Ctx) error {
	return nil
}
