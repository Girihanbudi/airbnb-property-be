package usecaseimpl

import (
	"airbnb-property-be/env/appcontext"
	module "airbnb-property-be/internal/app/property"
	errpreset "airbnb-property-be/internal/app/property/preset/error"
	"airbnb-property-be/internal/app/property/preset/request"
	transutil "airbnb-property-be/internal/app/translation/util"
	"airbnb-property-be/internal/pkg/stderror"
	"context"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (u Usecase) CreatePropertyType(ctx context.Context, cmd request.CreatePropertyType) (err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	fId, generateIdErr := gonanoid.New()
	if generateIdErr != nil {
		err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
		return
	}
	fName, uploadErr := u.Bucket.Upload(cmd.BFile, module.BucketGroupName, fId)
	if uploadErr != nil {
		err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
		return
	}

	propertyType := module.PropertyType{
		Code:       cmd.Code,
		LocaleCode: cmd.LocaleCode,
		Name:       cmd.Name,
		Link:       *fName,
	}

	createPropertyTypeErr := u.PropertyRepo.CreatePropertyType(ctx, &propertyType)
	if createPropertyTypeErr != nil {
		err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
		return
	}

	return
}
