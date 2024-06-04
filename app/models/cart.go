package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint
	ProductID uint
	Product   Product
	Quantity  int
}


type Purchase struct {
	gorm.Model
	UserID       uint
	TotalPrice   int
	Coupons      int
	PurchaseTime string
}
