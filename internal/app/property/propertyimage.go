package room

import "gorm.io/gorm"

type PropertyImage struct {
	gorm.Model
	PropertyId  string `json:"property_id" gorm:"not null"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Link        string `json:"link" gorm:"not null"`
}
