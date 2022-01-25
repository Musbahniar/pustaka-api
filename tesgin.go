package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "pong",
		// })
		c.SecureJSON(http.StatusOK, names)
	})
	router.Run(":8080")
}
