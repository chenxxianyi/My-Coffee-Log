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

	// Onboarding fields (v1)
	OnboardingCompleted  bool       `json:"onboarding_completed" gorm:"default:false"`
	PreferredLogMode     string     `json:"preferred_log_mode" gorm:"size:20;default:''"`     // quick | detailed
	PreferredCoffeeTypes string     `json:"preferred_coffee_types" gorm:"size:500;default:''"` // JSON array
	FirstRecordAt        *time.Time `json:"first_record_at"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
