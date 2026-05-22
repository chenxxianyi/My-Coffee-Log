package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Email        string         `json:"email" gorm:"uniqueIndex;size:255;not null"`
	Nickname     string         `json:"nickname" gorm:"size:100"`
	PasswordHash string         `json:"-" gorm:"size:255;not null"`
	AvatarURL    string         `json:"avatar_url" gorm:"size:500"`
	CoffeeLogs   []CoffeeLog    `json:"-" gorm:"foreignKey:UserID"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
