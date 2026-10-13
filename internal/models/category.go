package models

import (
	"gorm.io/gorm"
)

type Category struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Image string `json:"image"`
	gorm.Model
}
