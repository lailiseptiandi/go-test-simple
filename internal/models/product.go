package models

import "gorm.io/gorm"

type Product struct {
	ID          uint     `gorm:"primarykey" json:"id"`
	Name        string   `json:"name"`
	SKU         string   `json:"sku"`
	Slug        string   `json:"slug"`
	Image       string   `json:"image"`
	Description string   `json:"description"`
	CategoryID  int64    `json:"category_id"`
	Category    Category `gorm:"foreignKey:ID;references:CategoryID"`
	Price       float64  `json:"price"`
	gorm.Model
}
