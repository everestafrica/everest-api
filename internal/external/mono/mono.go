package mono

import (
	"github.com/gofiber/fiber"
	"os"
)

func VerifyWebhook() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		secret := os.Getenv("MONO_WEBHOOK_SECRET")
		webhookSecret := c.Get("mono-webhook-secret")
		if webhookSecret != secret {
			return c.Status(401).JSON(
				fiber.Map{
					"message": " Unauthorized user",
				})
		}
		return c.Next()
	}
}
