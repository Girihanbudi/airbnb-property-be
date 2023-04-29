package room

import (
	"github.com/Rhymond/go-money"
	"gorm.io/gorm"
)

type PropertyPrice struct {
	gorm.Model
	PropertyId   string      `json:"property_id" gorm:"not null"`
	Price        money.Money `json:"price" gorm:"embedded"`
	Unit         string      `json:"unit" gorm:"not null"`
	MinimumOrder int         `json:"minimum_order" gorm:"not null"`
}
