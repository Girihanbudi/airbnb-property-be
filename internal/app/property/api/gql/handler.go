package gql

import (
	"airbnb-property-be/internal/app/property/preset/request"
	"airbnb-property-be/internal/app/property/preset/response"
	"context"
)

func (h Handler) GetPropertyTypes(ctx context.Context, limit, page *int) (*response.GetPropertyTypes, error) {
	var req request.GetPropertyTypes
	if limit != nil {
		req.Pagination.Limit = *limit
	}

	if page != nil {
		req.Pagination.Page = *page
	}

	res, err := h.Property.GetPropertyTypes(ctx, req)
	if err != nil {
		return nil, err.Error
	}

	return &res, nil
}
