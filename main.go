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

type stock struct {
	symbol              string
	name                string
	sector              string
	price               float32
	price_per_earnings  sql.NullFloat64
	dividend_yield      float32
	earnings_per_share  float32
	fifty_two_week_low  float32
	fifty_two_week_high float32
	market_cap          float64
	EBITDA              float64
	price_per_sales     float32
	price_per_book      sql.NullFloat64
	sec_filings         string
}

func GetStocks(c *gin.Context) {
	db, err := GetDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM snp_500_financials")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		panic(err)
	}
	defer results.Close()

	var stocks []stock
	for results.Next() {
		var stock stock

		err := results.Scan(
			&stock.symbol,
			&stock.name,
			&stock.sector,
			&stock.price,
			&stock.price_per_earnings,
			&stock.dividend_yield,
			&stock.earnings_per_share,
			&stock.fifty_two_week_low,
			&stock.fifty_two_week_high,
			&stock.market_cap,
			&stock.EBITDA,
			&stock.price_per_sales,
			&stock.price_per_book,
			&stock.sec_filings,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err})
			panic(err)
		}

		stocks = append(stocks, stock)
	}

	c.JSON(http.StatusOK, gin.H{
		"stocks": stocks,
	})
}

func GetStock(c *gin.Context) {
	symbol := c.Param("symbol")

	// TODO: Get stock from database

	if len(symbol) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid Stock"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"stock": "stock",
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

	// Database
	db, err := GetDatabase()
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to PlanetScale!")
	defer db.Close()

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
