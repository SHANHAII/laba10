package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `json:"username" binding:"required,min=3"`
	Age      int    `json:"age" binding:"required,gte=18"`
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/user", func(c *gin.Context) {
		var u User
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "user validated", "data": u})
	})
	return r
}

func main() {
	SetupRouter().Run(":8081")
}