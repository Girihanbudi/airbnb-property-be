package wire

import (
	"airbnb-property-be/internal/pkg/aws"
	"airbnb-property-be/internal/pkg/aws/bucket"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	wire.Struct(new(aws.Options), "*"),
	aws.NewConstruct,
	wire.Struct(new(bucket.Options), "*"),
	bucket.NewBucket,
)
