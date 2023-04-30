package property

import "time"

type PropertyType struct {
	Code       string `json:"code" gorm:"primaryKey"`
	LocaleCode string `json:"locale_code" gorm:"primaryKey"`
	Name       string `json:"name" gorm:"not null"`
	Link       string `json:"link"`

	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
