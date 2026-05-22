package model

import (
	"time"

	"gorm.io/gorm"
)

type CoffeeLog struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id" gorm:"not null;index"`
	CoffeeName  string         `json:"coffee_name" gorm:"size:255"`
	CoffeeType  string         `json:"coffee_type" gorm:"size:50"`
	ShopName    string         `json:"shop_name" gorm:"size:255"`
	Location    string         `json:"location" gorm:"size:255"`
	ImageURL    string         `json:"image_url" gorm:"size:500"`
	DrinkDate   *time.Time     `json:"drink_date" gorm:"type:date"`
	Mood        string         `json:"mood" gorm:"size:50"`
	Notes       string         `json:"notes" gorm:"type:text"`
	Acidity     int            `json:"acidity" gorm:"default:0"`
	Bitterness  int            `json:"bitterness" gorm:"default:0"`
	Sweetness   int            `json:"sweetness" gorm:"default:0"`
	Body        int            `json:"body" gorm:"default:0"`
	Aroma       int            `json:"aroma" gorm:"default:0"`
	Aftertaste  int            `json:"aftertaste" gorm:"default:0"`
	AISummary   string         `json:"ai_summary" gorm:"type:text"`
	MoodTags    string         `json:"mood_tags" gorm:"type:varchar(500)"`
	SceneTags   string         `json:"scene_tags" gorm:"type:varchar(500)"`
	PairingTags string         `json:"pairing_tags" gorm:"type:varchar(500)"`
	BeanID      *uint          `json:"bean_id" gorm:"index"`
	BrewRatio   string         `json:"brew_ratio" gorm:"size:50"`
	WaterTemp   string         `json:"water_temp" gorm:"size:50"`
	GrindSize   string         `json:"grind_size" gorm:"size:50"`
	FlavorTags  []FlavorTag    `json:"flavor_tags" gorm:"many2many:coffee_log_flavor_tags"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	User User        `json:"-" gorm:"foreignKey:UserID"`
	Bean *CoffeeBean `json:"bean,omitempty" gorm:"foreignKey:BeanID"`
}
