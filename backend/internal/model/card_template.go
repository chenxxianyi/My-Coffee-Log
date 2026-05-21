package model

import "time"

type CardTemplate struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"size:100"`
	PreviewURL string    `json:"preview_url" gorm:"size:500"`
	LayoutType string    `json:"layout_type" gorm:"size:50"`
	IsPaid     bool      `json:"is_paid" gorm:"default:false"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
