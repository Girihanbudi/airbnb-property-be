package room

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type PropertyFacility struct {
	gorm.Model
	PropertyId  string `json:"property_id" gorm:"not null"`
	Name        string `json:"name" gorm:"not null"`
	Type        string `json:"type" gorm:"not null"`
	Description string `json:"description"`
	IconId      uint   `json:"icon_id"`

	Icon *PropertyFacilityIcon `json:"icon" gorm:"foreignKey:IconId"`
}

func (m *PropertyFacility) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), &m)
}

func (m PropertyFacility) Value() (driver.Value, error) {
	val, err := json.Marshal(m)
	return string(val), err
}
