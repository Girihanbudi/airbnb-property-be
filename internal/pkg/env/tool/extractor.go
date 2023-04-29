package tool

import (
	"airbnb-property-be/internal/pkg/env/config"
	gorm "airbnb-property-be/internal/pkg/gorm/config"
	httpServer "airbnb-property-be/internal/pkg/http/server/config"
	kafka "airbnb-property-be/internal/pkg/kafka/config"
	kafkaconsumer "airbnb-property-be/internal/pkg/kafka/consumer/config"
	kafkarouter "airbnb-property-be/internal/pkg/kafka/router/config"
)

func ExtractServerConfig(config config.Config) httpServer.Config {
	return config.HttpServer
}

func ExtractDBConfig(config config.Config) gorm.Config {
	return config.DB
}

func ExtractKafkaConfig(config config.Config) kafka.Config {
	return config.Kafka
}

func ExtractKafkaConsumerConfig(config config.Config) kafkaconsumer.Config {
	return config.Kafka.Consumer
}

func ExtractKafkaRouterConfig(config config.Config) kafkarouter.Config {
	return config.Kafka.Consumer.Router
}
