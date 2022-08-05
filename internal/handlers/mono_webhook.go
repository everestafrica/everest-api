package handlers

import (
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/gofiber/fiber/v2"
)

func VerifyMonoWebhook() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		secret := config.GetConf().MonoWebhookSecret
		webhookSecret := c.Get("mono-webhook-secret")
		if webhookSecret != secret {
			return c.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{
					"message": " Unauthorized user",
				})
		}
		c.Next()
		return nil
	}
}
