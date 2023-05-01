package request

import (
	"airbnb-property-be/internal/pkg/pagination"
	"airbnb-property-be/internal/pkg/validator"
)

type GetPropertyTypes struct {
	Pagination pagination.SQLPaging `json:"pagination"`
}

func (req *GetPropertyTypes) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
