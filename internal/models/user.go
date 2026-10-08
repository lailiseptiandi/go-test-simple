package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	ROLE_SUPER_ADMIN = 1
	ROLE_USER        = 2
)

type User struct {
	ID           uint   `gorm:"primarykey" json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username" gorm:"unique"`
	Password     string `json:"password"`
	Roles        uint   `json:"roles"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Email        string `json:"email" gorm:"unique"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
