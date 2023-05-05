package usecaseimpl

import (
	"airbnb-property-be/env/appcontext"
	module "airbnb-property-be/internal/app/property"
	errpreset "airbnb-property-be/internal/app/property/preset/error"
	"airbnb-property-be/internal/app/property/preset/request"
	"airbnb-property-be/internal/app/property/preset/response"
	transutil "airbnb-property-be/internal/app/translation/util"
	"airbnb-property-be/internal/pkg/stderror"
	"context"

	"github.com/thoas/go-funk"
)

func (u Usecase) GetPropertyTypes(ctx context.Context, cmd request.GetPropertyTypes) (res response.GetPropertyTypes, err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	propertyTypes, getPropertyTypesErr := u.PropertyRepo.GetPropertyTypes(ctx, &cmd.Pagination, clientLocale)
	if getPropertyTypesErr != nil {
		err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
		return
	}

	data := funk.Map(*propertyTypes, func(data module.PropertyType) response.PropertyType {
		var propertyType response.PropertyType

		propertyType.Code = data.Code
		propertyType.Name = data.Name
		propertyType.Link = data.Code
		propertyType.CreatedAt = data.CreatedAt
		propertyType.UpdatedAt = data.UpdatedAt

		return propertyType
	}).([]response.PropertyType)

	res.PropertyTypes = &data
	res.Pagination = &cmd.Pagination

	return
}
