package wire

import (
	"airbnb-property-be/internal/pkg/http/server"
	"airbnb-property-be/internal/pkg/http/server/router"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	router.NewRouter,

	wire.Struct(new(server.Options), "*"),
	server.NewServer,
)
