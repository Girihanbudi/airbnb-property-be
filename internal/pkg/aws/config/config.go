package config

import (
	bucket "airbnb-property-be/internal/pkg/aws/bucket/config"
)

type Config struct {
	Bucket bucket.Config `mapstructure:"bucket"`
}
