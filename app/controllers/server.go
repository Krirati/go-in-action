package controllers

import (
	"learning/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
    Config config.Config
}

func New(config *config.Config) *Server {
	router := gin.Default()
	server := &Server{
		Router: router,
        Config: *config,
	}
    server.SetupRoutes()
	return server
}

func (s *Server) Run() {
	s.Router.Run(s.Config.ServerAddress)
}
