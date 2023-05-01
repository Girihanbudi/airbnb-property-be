package usecaseimpl

import (
	"airbnb-property-be/internal/app/property/repo"
	"airbnb-property-be/internal/pkg/aws/bucket"
)

type Options struct {
	PropertyRepo repo.IProperty
	Bucket       *bucket.Engine
}

type Usecase struct {
	Options
}

func NewPropertyUsecase(options Options) *Usecase {
	return &Usecase{options}
}
