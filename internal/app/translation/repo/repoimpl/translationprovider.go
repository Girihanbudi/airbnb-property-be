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

var TranslationRepo Repo

func NewTranslationRepo(options Options) *Repo {
	TranslationRepo = Repo{options}
	return &TranslationRepo
}
