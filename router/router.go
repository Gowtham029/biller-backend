package router

import (
	"identity/handler"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("identity service is running")
	})
	// Docs
	app.Get("/docs/*", swagger.HandlerDefault)

	// Middleware
	api := app.Group("/api", logger.New())

	api.Post("/login", handler.Login)

	api.Post("/register", handler.CreateUser)
}
