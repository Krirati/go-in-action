package controllers

import (
	"learning/app/models"
	"learning/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	Config config.Config
	Store  models.Store
}

func New(config *config.Config, store models.Store) *Server {
	router := gin.Default()
	server := &Server{
		Router: router,
		Config: *config,
		Store:  store,
	}
	server.SetupRoutes()
	return server
}

func (s *Server) Run() {
	s.Router.Run(s.Config.ServerAddress)
}
