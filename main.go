package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	return db, err
}

func GetStocks(c *gin.Context) {
	// TODO: Get stocks from database

	c.JSON(http.StatusOK, gin.H{
		"message": "get users",
	})
}

func GetStock(c *gin.Context) {
	symbol := c.Param("symbol")

	// TODO: Get stock from database

	if len(symbol) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Stock"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"symbol": symbol,
	})

}

func CreateStock(c *gin.Context) {
	// TODO: Create stock in database

	c.JSON(http.StatusCreated, gin.H{
		"message": "Crater Created",
	})
}

func UpdateStock(c *gin.Context) {
	symbol := c.Param("symbol")

	// TODO: Update stock in database

	c.JSON(http.StatusOK, gin.H{
		"symbol":  symbol,
		"message": "Crater Updated",
	})
}

func DeleteStock(c *gin.Context) {
	symbol := c.Param("symbol")

	// TODO: Delete stock from database

	c.JSON(http.StatusOK, gin.H{
		"symbol":  symbol,
		"message": "Crater Updated",
	})
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := gin.Default()

	db, err := GetDatabase()
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to PlanetScale!")

	r.GET("/stocks/", GetStocks)
	r.GET("/stocks/:symbol", GetStock)
	r.POST("/stocks/", CreateStock)
	r.PUT("/stocks/:symbol", UpdateStock)
	r.DELETE("/stocks/:symbol", DeleteStock)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
