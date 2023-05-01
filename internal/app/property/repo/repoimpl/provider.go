package repoimpl

import (
	"airbnb-property-be/internal/pkg/gorm"
)

type Options struct {
	Gorm *gorm.Engine
}

type Repo struct {
	Options
}

func NewPropertyRepo(options Options) *Repo {
	return &Repo{options}
}
