package routes

import (
	"simple-web-cart/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", nil)
	})
	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("register", nil)
	})
	app.Post("/login", controllers.Login)
	app.Post("/register", controllers.Register)

	// Routes for Products CRUD
	app.Get("/products", controllers.GetProducts)
	app.Post("/products", controllers.CreateProduct)
	app.Put("/products/:id", controllers.UpdateProduct)
	app.Delete("/products/:id", controllers.DeleteProduct)

	app.Get("/cart", controllers.GetCart)
	app.Post("/cart", controllers.AddToCart)
	app.Delete("/cart/:id", controllers.RemoveFromCart)
	app.Get("/cart/checkout", controllers.Checkout)

	app.Get("/cart", func(c *fiber.Ctx) error {
		return c.Render("cart", nil)
	})

	app.Get("/summary", func(c *fiber.Ctx) error {
		return c.Render("summary", nil)
	})

	app.Get("history", controllers.PurchaseHistory)
}
