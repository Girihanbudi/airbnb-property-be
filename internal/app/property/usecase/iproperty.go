package usecase

import (
	"airbnb-property-be/internal/app/property/preset/request"
	"airbnb-property-be/internal/app/property/preset/response"
	"airbnb-property-be/internal/pkg/stderror"
	"context"
)

type IProperty interface {
	CreatePropertyType(ctx context.Context, cmd request.CreatePropertyType) (err *stderror.StdError)
	GetPropertyTypes(ctx context.Context, cmd request.GetPropertyTypes) (res response.GetPropertyTypes, err *stderror.StdError)
}
