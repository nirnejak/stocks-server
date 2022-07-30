package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")

		if len(name) < 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name should be longer than 1 character"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"timeInQueue": "Hello " + name})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
