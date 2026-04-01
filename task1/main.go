package main

import (
	"github.com/gin-gonic/gin"
	"net/http" 
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}

func main() {
	r := SetupRouter()
	r.Run(":8080")
}