package response

import (
	"airbnb-property-be/internal/pkg/pagination"
)

type GetPropertyTypes struct {
	PropertyTypes *[]PropertyType       `json:"propertyTypes"`
	Pagination    *pagination.SQLPaging `json:"pagination"`
}
