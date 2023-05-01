//go:build wireinject
// +build wireinject

package app

import (
	aws "airbnb-property-be/internal/pkg/aws/wire"
	env "airbnb-property-be/internal/pkg/env/wire"
	gorm "airbnb-property-be/internal/pkg/gorm/wire"
	http "airbnb-property-be/internal/pkg/http/server/wire"
	kafka "airbnb-property-be/internal/pkg/kafka/wire"

	property "airbnb-property-be/internal/app/property/wire"
	translation "airbnb-property-be/internal/app/translation/wire"

	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	wire.Struct(new(Options), "*"),
	wire.Struct(new(App), "*"),
)

func NewApp() (*App, error) {
	panic(
		wire.Build(
			env.PackageSet,
			gorm.PackageSet,
			http.PackageSet,
			kafka.PackageSet,
			aws.PackageSet,

			AppSet,

			translation.ModuleSet,
			property.ModuleSet,
		),
	)
}
