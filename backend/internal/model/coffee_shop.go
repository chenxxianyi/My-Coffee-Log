package model

import (
	"time"

	"gorm.io/gorm"
)

type CoffeeShop struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id" gorm:"not null;index"`
	Name        string         `json:"name" gorm:"size:255;not null"`
	Address     string         `json:"address" gorm:"size:500"`
	Rating      int            `json:"rating" gorm:"default:0"`
	ImageURL    string         `json:"image_url" gorm:"size:500"`
	VisitCount  int            `json:"visit_count" gorm:"default:0"`
	LastVisitAt *time.Time     `json:"last_visit_at" gorm:"type:date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	User User `json:"-" gorm:"foreignKey:UserID"`
}
