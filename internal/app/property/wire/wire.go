package wire

import (
	"airbnb-property-be/internal/app/property/api/gql"
	"airbnb-property-be/internal/app/property/api/rest"
	"airbnb-property-be/internal/app/property/repo"
	"airbnb-property-be/internal/app/property/repo/repoimpl"
	"airbnb-property-be/internal/app/property/usecase"
	"airbnb-property-be/internal/app/property/usecase/usecaseimpl"

	"github.com/google/wire"
)

var ModuleSet = wire.NewSet(
	repoSet,
	usecaseSet,
	apiSet,
)

var repoSet = wire.NewSet(
	wire.Struct(new(repoimpl.Options), "*"),
	repoimpl.NewPropertyRepo,
	wire.Bind(new(repo.IProperty), new(*repoimpl.Repo)),
)

var usecaseSet = wire.NewSet(
	wire.Struct(new(usecaseimpl.Options), "*"),
	usecaseimpl.NewPropertyUsecase,
	wire.Bind(new(usecase.IProperty), new(*usecaseimpl.Usecase)),
)

var apiSet = wire.NewSet(
	wire.Struct(new(gql.Options), "*"),
	gql.NewPropertyHandler,

	wire.Struct(new(rest.Options), "*"),
	rest.NewPropertyHandler,
)
