package model

import "time"

type FlavorTag struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"uniqueIndex;size:100"`
	Label     string    `json:"label" gorm:"size:100"`
	Color     string    `json:"color" gorm:"size:50"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
