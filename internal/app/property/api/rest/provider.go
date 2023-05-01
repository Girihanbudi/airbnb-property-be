package rest

import (
	"github.com/gin-gonic/gin"

	propertyusecase "airbnb-property-be/internal/app/property/usecase"
)

type Options struct {
	Router *gin.Engine

	Property propertyusecase.IProperty
}

type Handler struct {
	Options
}

func NewPropertyHandler(options Options) *Handler {
	return &Handler{options}
}
