package room

import (
	"gorm.io/gorm"
)

type PropertyFacilityIcon struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
	Link string `json:"link" gorm:"not null"`
}
