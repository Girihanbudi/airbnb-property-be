package config

import (
	bucket "airbnb-property-be/internal/pkg/aws/bucket/config"
)

type Config struct {
	AccessKey string        `mapstructure:"accesskey"`
	SecretKey string        `mapstructure:"secretkey"`
	Region    string        `mapstructure:"region"`
	Url       string        `mapstructure:"url"`
	Bucket    bucket.Config `mapstructure:"bucket"`
}
