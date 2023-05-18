package config

import (
	aws "airbnb-property-be/internal/pkg/aws/config"
	cache "airbnb-property-be/internal/pkg/cache/config"
	elastic "airbnb-property-be/internal/pkg/elasticsearch/config"
	gorm "airbnb-property-be/internal/pkg/gorm/config"
	httpserver "airbnb-property-be/internal/pkg/http/server/config"
	jwt "airbnb-property-be/internal/pkg/jwt/config"
	kafka "airbnb-property-be/internal/pkg/kafka/config"
)

type Config struct {
	Origins    []string          `mapstructure:"origins"`
	Stage      string            `mapstructure:"stage"`
	Domain     string            `mapstructure:"domain"`
	HttpServer httpserver.Config `mapstructure:"httpserver"`
	DB         gorm.Config       `mapstructure:"db"`
	Jwt        jwt.Config        `mapstructure:"jwt"`
	Cache      cache.Config      `mapstructure:"cache"`
	Kafka      kafka.Config      `mapstructure:"kafka"`
	Aws        aws.Config        `mapstructure:"aws"`
	Elastic    elastic.Config    `mapstructure:"elastic"`
}
