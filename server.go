package main

import (
	"errors"
	"fmt"
	"github.com/everestafrica/everest-api/internal/commons/constants"
	"github.com/everestafrica/everest-api/internal/config"
	"github.com/everestafrica/everest-api/internal/handlers"
	"github.com/everestafrica/everest-api/internal/route"
	"github.com/everestafrica/everest-api/internal/scheduler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type server struct {
	cfg            *config.Config
	requestTimeout time.Duration
}

func (s *server) Start() error {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	app := fiber.New()

	app.Use(requestid.New(requestid.Config{
		Header: constants.RequestIdentifier,
		Generator: func() string {
			return uuid.NewString()
		},
	}))
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("You're home, yaay!!")
	})

	app.Static("/", "./public")

	route.RegisterRoutes(app)

	setupSystemRouteHandler(app)

	scheduler.RegisterSchedulers()

	go func() {
		if err := app.Listen(":" + s.cfg.Port); err != nil && err != http.ErrServerClosed {
			log.Fatal(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), s.requestTimeout*time.Second)

	defer cancel()

	if err := app.Shutdown(); err != nil {
		return errors.New(fmt.Sprintf("Server forced to shutdown: %v", err))

	}

	log.Println("Server exiting")

	return nil

}

func setupSystemRouteHandler(app *fiber.App) {
	app.Use(cors.New(handlers.Cors()))

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
	// 405 Handler
	//router.NoMethod(handlers.Http405Handler())
}
