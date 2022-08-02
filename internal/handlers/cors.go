package handlers

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors() cors.Config {
	return cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowMethods:     "POST, OPTIONS, GET, PUT, PATCH",
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
	}
}
