package controllers

import (
	"simple-web-cart/app/database"
	"simple-web-cart/app/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PurchaseHistory(c *fiber.Ctx) error {
	userID, _ := c.Locals("userID").(uint)

	var purchases []models.Purchase
	result := database.DB.Where("user_id = ?", userID).Find(&purchases)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	type PurchaseView struct {
		models.Purchase
		Status string
	}

	var purchaseViews []PurchaseView
	for _, purchase := range purchases {
		purchaseTime, _ := time.Parse("2006-01-02 15:04:05", purchase.PurchaseTime)
		status := "open"
		if time.Since(purchaseTime).Hours() > 3 {
			status = "closed"
		}
		purchaseViews = append(purchaseViews, PurchaseView{purchase, status})
	}

	return c.Render("history", fiber.Map{
		"Purchases": purchaseViews,
	})
}
