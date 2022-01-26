package controllers

import "github.com/gin-gonic/gin"

func (s *Server) initializeRoutes() {
	v0 := s.Router.Group("/")
	{
		v0.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": true, "message": "Server Pustaka API"})
		})
	}

	// 	v1 := s.Router.Group("/api/v1")
	// 	{

	// 	}
}
