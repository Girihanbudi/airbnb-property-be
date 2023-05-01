package repoimpl

import (
	module "airbnb-property-be/internal/app/property"
	"context"
)

func (r Repo) CreatePropertyType(ctx context.Context, propertyType *module.PropertyType) (err error) {
	err = r.Gorm.DB.Create(propertyType).Error
	return
}
