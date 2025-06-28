package controllers

func (s *Server) SetupRoutes() {
	s.Router.GET("/", s.GetHealth)
}
