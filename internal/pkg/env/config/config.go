package config

import (
	aws "airbnb-property-be/internal/pkg/aws/config"
	cache "airbnb-property-be/internal/pkg/cache/config"
	credential "airbnb-property-be/internal/pkg/credential/config"
	elastic "airbnb-property-be/internal/pkg/elasticsearch/config"
	gorm "airbnb-property-be/internal/pkg/gorm/config"
	httpserver "airbnb-property-be/internal/pkg/http/server/config"
	jwt "airbnb-property-be/internal/pkg/jwt/config"
	kafka "airbnb-property-be/internal/pkg/kafka/config"
)

type Config struct {
	Stage      string            `mapstructure:"stage"`
	Origins    []string          `mapstructure:"origins"`
	Domain     string            `mapstructure:"domain"`
	Creds      credential.Config `mapstructure:"creds"`
	HttpServer httpserver.Config `mapstructure:"httpserver"`
	DB         gorm.Config       `mapstructure:"db"`
	Jwt        jwt.Config        `mapstructure:"jwt"`
	Cache      cache.Config      `mapstructure:"cache"`
	Elastic    elastic.Config    `mapstructure:"elastic"`
	Aws        aws.Config        `mapstructure:"aws"`
	Kafka      kafka.Config      `mapstructure:"kafka"`
}
