package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/stocks", GetStocks)
	r.GET("/stocks/:symbol", GetStock)
	r.POST("/stocks", CreateStock)
	r.PUT("/stocks/:symbol", UpdateStock)
	r.DELETE("/stocks/:symbol", DeleteStock)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
