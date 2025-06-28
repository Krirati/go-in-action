package controllers

import "github.com/gin-gonic/gin"

type Server struct {
	Router *gin.Engine
}

func New() *Server {
	router := gin.Default()
	server := &Server{
		Router: router,
	}
    server.SetupRoutes()
	return server
}

func (s *Server) Run() {
	s.Router.Run()
}
