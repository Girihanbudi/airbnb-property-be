package wire

import (
	"airbnb-property-be/internal/pkg/env"
	"airbnb-property-be/internal/pkg/env/tool"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	env.ProvideEnv,
	tool.ExtractServerConfig,
	tool.ExtractDBConfig,
	tool.ExtractKafkaConfig,
	tool.ExtractKafkaConsumerConfig,
	tool.ExtractKafkaRouterConfig,
)
