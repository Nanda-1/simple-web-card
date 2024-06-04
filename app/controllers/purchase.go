package controllers

import (
	"simple-web-cart/app/database"
	"simple-web-cart/app/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Checkout(c *fiber.Ctx) error {
	userID, _ := c.Locals("userID").(uint)

	var cartItems []models.Cart
	result := database.DB.Where("user_id = ?", userID).Find(&cartItems)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	var totalPrice int
	var coupons int
	var products []fiber.Map
	for _, item := range cartItems {
		var product models.Product
		database.DB.First(&product, item.ProductID)
		totalPrice += product.Price * item.Quantity
		if product.Price > 50000 {
			coupons += 1
		}
		products = append(products, fiber.Map{
			"Name":     product.Name,
			"Price":    product.Price,
			"Quantity": item.Quantity,
			"Subtotal": product.Price * item.Quantity,
		})
	}
	coupons += totalPrice / 100000

	purchaseTime := time.Now().Format("2006-01-02 15:04:05")
	purchase := models.Purchase{UserID: userID, TotalPrice: totalPrice, Coupons: coupons, PurchaseTime: purchaseTime}
	result = database.DB.Create(&purchase)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	database.DB.Where("user_id = ?", userID).Delete(&models.Cart{})

	return c.Render("summary", fiber.Map{
		"Products":   products,
		"TotalPrice": totalPrice,
		"Coupons":    coupons,
	})
}
