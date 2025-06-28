package controllers

import "github.com/gin-gonic/gin"

func (s *Server) GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "healthy",
	})
}
