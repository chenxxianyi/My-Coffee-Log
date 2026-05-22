package model

import (
	"time"

	"gorm.io/gorm"
)

type CoffeeBean struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	UserID          uint           `json:"user_id" gorm:"not null;index"`
	Name            string         `json:"name" gorm:"size:255;not null"`
	Origin          string         `json:"origin" gorm:"size:255"`
	ProcessingMethod string        `json:"processing_method" gorm:"size:100"`
	RoastLevel      string         `json:"roast_level" gorm:"size:50"`
	Roaster         string         `json:"roaster" gorm:"size:255"`
	ImageURL        string         `json:"image_url" gorm:"size:500"`
	UsageCount      int            `json:"usage_count" gorm:"default:0"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	User User `json:"-" gorm:"foreignKey:UserID"`
}
