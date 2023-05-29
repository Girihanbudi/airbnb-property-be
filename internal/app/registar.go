package app

import (
	_ "airbnb-property-be/docs"
	"airbnb-property-be/graph"
	gqlproperty "airbnb-property-be/internal/app/property/api/gql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

// Defining the Graphql handler
func graphqlHandler(propertyHandler gqlproperty.Handler) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{Resolvers: &graph.Resolver{
			Property: propertyHandler,
		}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (a App) registerHttpHandler() {
	// Register rest handler
	a.PropertyRestHandler.RegisterApi()

	// Register modules to graph solver handler
	a.HttpServer.Router.GET("/graph", graphqlHandler(
		*a.PropertyGqlHandler,
	))

	// Register swagger documentation
	a.HttpServer.Router.GET("/docs/*any", ginswagger.WrapHandler(swaggerfiles.Handler))
}
