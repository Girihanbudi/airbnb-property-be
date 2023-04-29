package room

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Property struct {
	Id          string             `json:"id" gorm:"primaryKey"`
	Name        string             `json:"name" gorm:"not null"`
	Description string             `json:"description"`
	Rates       float32            `json:"rates"`
	Likes       int                `json:"likes"`
	Types       pq.StringArray     `json:"types" gorm:"type:text[]"`
	Facilities  []PropertyFacility `json:"facilities" gorm:"foreignKey:PropertyId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Prices      []PropertyPrice    `json:"prices" gorm:"foreignKey:PropertyId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Images      []PropertyImage    `json:"images" gorm:"foreignKey:PropertyId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	CreatedAt time.Time      `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
