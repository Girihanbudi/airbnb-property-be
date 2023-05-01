package graph

import (
	property "airbnb-property-be/internal/app/property/api/gql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Property property.Handler
}
