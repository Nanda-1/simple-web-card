package main

import (
	"simple-web-cart/app/database"
	"simple-web-cart/app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	database.InitDB()

	// Set up Fiber with HTML templates
	engine := html.New("./app/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	app.Listen(":3000")
}
