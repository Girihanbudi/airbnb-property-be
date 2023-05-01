package repoimpl

import (
	module "airbnb-property-be/internal/app/property"
	"airbnb-property-be/internal/pkg/pagination"
	"context"
)

func (r Repo) GetPropertyTypes(ctx context.Context, paging *pagination.SQLPaging, localeCode string) (propertyTypes *[]module.PropertyType, err error) {
	query := r.Gorm.DB.Where("locale_code = ?", localeCode)

	err = query.
		Scopes(pagination.GormPaginate(&module.PropertyType{}, paging, query)).
		Find(&propertyTypes).Error

	return
}
