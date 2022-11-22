package main

import (
	"identity/database"
	_ "identity/docs"
	"identity/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(requestid.New())

	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	database.ConnectDB()
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":9090"))
}
