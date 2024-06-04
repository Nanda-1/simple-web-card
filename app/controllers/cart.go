package controllers

import (
	"simple-web-cart/app/database"
	"simple-web-cart/app/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func AddToCart(c *fiber.Ctx) error {
	userID, _ := c.Locals("userID").(uint)
	productID, err := strconv.Atoi(c.FormValue("product_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid product ID")
	}
	quantity, err := strconv.Atoi(c.FormValue("quantity"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid quantity")
	}

	cart := models.Cart{UserID: userID, ProductID: uint(productID), Quantity: quantity}
	result := database.DB.Create(&cart)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.Redirect("/history")
}

// controllers/cart.go
func GetCart(c *fiber.Ctx) error {
	userID, _ := c.Locals("userID").(uint)

	var cartItems []models.Cart
	result := database.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.Render("cart", fiber.Map{
		"CartItems": cartItems,
	})
}

// controllers/cart.go
func RemoveFromCart(c *fiber.Ctx) error {
	userID, _ := c.Locals("userID").(uint)
	cartID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid cart ID")
	}

	var cart models.Cart
	result := database.DB.Where("id = ? AND user_id = ?", cartID, userID).Delete(&cart)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.Redirect("/cart")
}
