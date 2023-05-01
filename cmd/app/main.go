package main

import (
	"airbnb-property-be/internal/app"
	"airbnb-property-be/internal/pkg/env"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title           Airbnb Property Backend API
// @version         1.0
// @description     Airbnb Property Backend Service API
// @termsOfService  https://airbnb.co.id

// @contact.name   API Support
// @contact.url    https://airbnb.co.id/support
// @contact.email  support@airbnb.co.id

// @host      localhost/api
// @BasePath  /property

// @securityDefinitions.basic BasicAuth
func main() {
	// init app environment
	defaultEnvOps := env.NewDefaultOptions()
	env.InitEnv(defaultEnvOps)

	// create app context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// create service app
	serviceApp, err := app.NewApp()
	if err != nil {
		log.Panic(err)
	}

	serviceApp.Run(ctx)
	stop()
}
