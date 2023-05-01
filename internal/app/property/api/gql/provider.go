package gql

import (
	"airbnb-property-be/internal/app/property/usecase"
)

type Options struct {
	Property usecase.IProperty
}

type Handler struct {
	Options
}

func NewPropertyHandler(options Options) *Handler {
	return &Handler{options}
}
