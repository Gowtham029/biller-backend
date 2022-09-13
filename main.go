package main

import (
	"biller/database"
	_ "biller/docs"
	"biller/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":9090"))
}
