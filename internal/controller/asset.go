package controller

import (
	"github.com/everestafrica/everest-api/internal/handlers"
	"github.com/everestafrica/everest-api/internal/service"
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
	assetService service.IAssetService
}

// NewAssetController instantiates assetController
func NewAssetController() IAssetController {
	return &assetController{
		assetService: service.NewAssetService(),
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
