package server

import (
	"airbnb-property-be/internal/pkg/credential"
	"airbnb-property-be/internal/pkg/http/server/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

const Instance = "HTTP Server"

type Options struct {
	config.Config
	Router *gin.Engine
	Creds  credential.TlsCredentials
}

type Server struct {
	Options
	address string
	server  *http.Server
}

func NewServer(options Options) *Server {
	return &Server{
		Options: options,
		address: fmt.Sprintf("%s:%s", options.Host, options.Port)
	}
}
