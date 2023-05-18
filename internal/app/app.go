package app

import (
	authmid "airbnb-property-be/internal/app/middleware/auth"
	cookiemid "airbnb-property-be/internal/app/middleware/cookie"
	elasticmid "airbnb-property-be/internal/app/middleware/elastic"
	"airbnb-property-be/internal/pkg/cache/auth"
	"airbnb-property-be/internal/pkg/cache/otp"
	elastic "airbnb-property-be/internal/pkg/elasticsearch"
	"airbnb-property-be/internal/pkg/http/server"
	httprouter "airbnb-property-be/internal/pkg/http/server/router"
	kafkaconsumer "airbnb-property-be/internal/pkg/kafka/consumer"
	kafkaproducer "airbnb-property-be/internal/pkg/kafka/producer"
	"airbnb-property-be/internal/pkg/log"
	"context"
	"sync"

	"github.com/gin-gonic/gin"

	propertygql "airbnb-property-be/internal/app/property/api/gql"
	propertyrest "airbnb-property-be/internal/app/property/api/rest"
	translation "airbnb-property-be/internal/app/translation/repo"
)

var Instance = "App"

type Options struct {
	HttpServer    *server.Server
	EventListener *kafkaconsumer.Listener
	EventProducer *kafkaproducer.Producer

	Translation         translation.ITranslation
	PropertyGqlHandler  *propertygql.Handler
	PropertyRestHandler *propertyrest.Handler
}

type App struct {
	Options
}

// Run all the modules of the app.
func (a App) Run(ctx context.Context) {
	a.runModules(ctx)
	a.stopModules()
}

func (a App) runModules(ctx context.Context) {
	log.Event(Instance, "Starting...")

	// init app cache
	auth.InitAuthCache()
	otp.InitOtpCache()

	// Init elasticsearch client
	elastic.InitElasticSearch()

	// Create required index in elastic
	elasticmid.CreateIndex()

	// recover from panic
	a.HttpServer.Router.Use(gin.Recovery())

	// GIN apply CORS setting
	a.HttpServer.Router.Use(httprouter.DefaultCORSSetting())

	// GIN log request and response to elastic
	a.HttpServer.Router.Use(elasticmid.LogRequestToElastic())

	// GIN bind all cookie
	a.HttpServer.Router.Use(cookiemid.BindAll())

	// GIN bind access token if any
	// bind access token in all route to adapt with graphql endpoint
	a.HttpServer.Router.Use(authmid.GinBindAccessToken())

	// Register all routes
	a.registerHttpHandler()

	go func() {
		err := a.HttpServer.Start()
		if err != nil {
			log.Fatal(Instance, "failed to start http server", err)
		}
	}()

	<-ctx.Done()
}

func (a App) stopModules() {
	log.Event(Instance, "Stoping...")

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		err := a.EventProducer.Stop()
		if err != nil {
			log.Fatal(Instance, "failed to stop event producer", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := a.HttpServer.Stop()
		if err != nil {
			log.Fatal(Instance, "failed to stop http server", err)
		}
	}()

	wg.Wait()
}
