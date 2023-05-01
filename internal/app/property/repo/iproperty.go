package repo

import (
	module "airbnb-property-be/internal/app/property"
	"airbnb-property-be/internal/pkg/pagination"
	"context"
)

type IProperty interface {
	CreatePropertyType(ctx context.Context, propertyType *module.PropertyType) (err error)
	GetPropertyTypes(ctx context.Context, paging *pagination.SQLPaging, localeCode string) (propertyTypes *[]module.PropertyType, err error)
}
