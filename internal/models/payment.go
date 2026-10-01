package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID             uint    `gorm:"primarykey" json:"id"`
	OrderID        string  `json:"order_id"`
	Amount         float64 `json:"amount"`
	IdempotencyKey string  `json:"-" gorm:"not null;uniqueIndex"`
	Status         string  `json:"status" gorm:"default:pending"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
