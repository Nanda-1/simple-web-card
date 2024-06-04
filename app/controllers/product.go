package controllers

import (
	"simple-web-cart/app/database"
	"simple-web-cart/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	result := database.DB.Find(&products)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.Render("product", fiber.Map{
		"Products": products,
	})
}

// CreateProduct membuat produk baru.
func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	result := database.DB.Create(&product)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.JSON(product)
}

// UpdateProduct memperbarui produk yang ada.
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	result := database.DB.Model(&models.Product{}).Where("id = ?", id).Updates(&product)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.JSON(product)
}

// DeleteProduct menghapus produk yang ada.
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	result := database.DB.Delete(&models.Product{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.SendString("Product successfully deleted")
}
