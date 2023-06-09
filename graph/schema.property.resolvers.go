package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"airbnb-property-be/graph/model"
	"airbnb-property-be/internal/app/property/preset/response"
	"context"

	"github.com/thoas/go-funk"
)

// PropertyTypes is the resolver for the propertyTypes field.
func (r *queryResolver) PropertyTypes(ctx context.Context, limit *int, page *int) (*model.PropertyTypes, error) {
	data, err := r.Resolver.Property.GetPropertyTypes(ctx, limit, page)
	if err != nil {
		return nil, err
	}

	propertyTypes := funk.Map(*data.PropertyTypes, func(data response.PropertyType) *model.PropertyType {
		var propertyType model.PropertyType

		propertyType.Code = data.Code
		propertyType.Link = &data.Link
		propertyType.Name = data.Name
		propertyType.CreatedAt = data.CreatedAt
		propertyType.UpdatedAt = data.UpdatedAt

		return &propertyType
	}).([]*model.PropertyType)

	pagination := data.Pagination
	meta := model.Pagination{
		Limit:    &pagination.Limit,
		Page:     &pagination.Page,
		PageSize: &pagination.PageSize,
	}

	response := model.PropertyTypes{
		Data: propertyTypes,
		Meta: &meta,
	}

	return &response, nil
}
