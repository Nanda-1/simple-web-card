package database

import (
	"log"
	"simple-web-cart/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=admin dbname=simple_web_cart port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Cart{}, &models.Purchase{})
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}
}
