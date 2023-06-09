// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"airbnb-property-be/internal/app/property/api/gql"
	"airbnb-property-be/internal/app/property/api/rest"
	repoimpl2 "airbnb-property-be/internal/app/property/repo/repoimpl"
	"airbnb-property-be/internal/app/property/usecase/usecaseimpl"
	"airbnb-property-be/internal/app/translation/repo/repoimpl"
	"airbnb-property-be/internal/pkg/aws"
	"airbnb-property-be/internal/pkg/aws/bucket"
	"airbnb-property-be/internal/pkg/env"
	"airbnb-property-be/internal/pkg/env/tool"
	"airbnb-property-be/internal/pkg/gorm"
	"airbnb-property-be/internal/pkg/http/server"
	"airbnb-property-be/internal/pkg/http/server/router"
	"airbnb-property-be/internal/pkg/kafka"
	"airbnb-property-be/internal/pkg/kafka/consumer"
	"airbnb-property-be/internal/pkg/kafka/producer"
	router2 "airbnb-property-be/internal/pkg/kafka/router"
	"github.com/google/wire"
)

import (
	_ "airbnb-property-be/docs"
)

// Injectors from wire.go:

func NewApp() (*App, error) {
	config := env.ProvideEnv()
	configConfig := tool.ExtractServerConfig(config)
	engine := router.NewRouter()
	options := server.Options{
		Config: configConfig,
		Router: engine,
	}
	serverServer := server.NewServer(options)
	config2 := tool.ExtractKafkaConsumerConfig(config)
	config3 := tool.ExtractKafkaConfig(config)
	config4 := tool.ExtractKafkaRouterConfig(config)
	routerOptions := router2.Options{
		Config: config4,
	}
	routerRouter := router2.NewRouter(routerOptions)
	kafkaOptions := kafka.Options{
		Config: config3,
		Router: routerRouter,
	}
	client := kafka.NewSaramaClient(kafkaOptions)
	consumerOptions := consumer.Options{
		Config: config2,
		Client: client,
		Router: routerRouter,
	}
	listener := consumer.NewEventListener(consumerOptions)
	producerOptions := producer.Options{
		Client: client,
	}
	producerProducer := producer.NewEventProducer(producerOptions)
	config5 := tool.ExtractDBConfig(config)
	gormOptions := gorm.Options{
		Config: config5,
	}
	gormEngine := gorm.NewORM(gormOptions)
	repoimplOptions := repoimpl.Options{
		Gorm: gormEngine,
	}
	repo := repoimpl.NewTranslationRepo(repoimplOptions)
	options2 := repoimpl2.Options{
		Gorm: gormEngine,
	}
	repoimplRepo := repoimpl2.NewPropertyRepo(options2)
	config6 := tool.ExtractAwsBucketConfig(config)
	config7 := tool.ExtractAwsConfig(config)
	awsOptions := aws.Options{
		Config: config7,
	}
	awsAws := aws.NewAwsClient(awsOptions)
	bucketOptions := bucket.Options{
		Config:    config6,
		AwsClient: awsAws,
	}
	bucketEngine := bucket.NewBucket(bucketOptions)
	usecaseimplOptions := usecaseimpl.Options{
		PropertyRepo: repoimplRepo,
		Bucket:       bucketEngine,
	}
	usecase := usecaseimpl.NewPropertyUsecase(usecaseimplOptions)
	gqlOptions := gql.Options{
		Property: usecase,
	}
	handler := gql.NewPropertyHandler(gqlOptions)
	restOptions := rest.Options{
		Router:   engine,
		Property: usecase,
	}
	restHandler := rest.NewPropertyHandler(restOptions)
	appOptions := Options{
		HttpServer:          serverServer,
		EventListener:       listener,
		EventProducer:       producerProducer,
		Translation:         repo,
		PropertyGqlHandler:  handler,
		PropertyRestHandler: restHandler,
	}
	app := &App{
		Options: appOptions,
	}
	return app, nil
}

// wire.go:

var AppSet = wire.NewSet(wire.Struct(new(Options), "*"), wire.Struct(new(App), "*"))
