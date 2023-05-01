package request

import (
	"airbnb-property-be/internal/pkg/validator"
	"mime/multipart"
)

type CreatePropertyType struct {
	Code       string                `form:"code" validation:"required"`
	LocaleCode string                `form:"localeCode" validation:"required"`
	Name       string                `form:"name" validation:"required"`
	File       *multipart.FileHeader `form:"file"`
	BFile      []byte
}

func (req *CreatePropertyType) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
