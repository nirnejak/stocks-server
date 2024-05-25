package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	return db, err
}

type STOCK struct {
	symbol              sql.NullString
	name                sql.NullString
	sector              sql.NullString
	price               sql.NullFloat64
	price_per_earnings  sql.NullFloat64
	dividend_yield      sql.NullFloat64
	earnings_per_share  sql.NullFloat64
	fifty_two_week_low  sql.NullFloat64
	fifty_two_week_high sql.NullFloat64
	market_cap          sql.NullFloat64
	EBITDA              sql.NullFloat64
	price_per_sales     sql.NullFloat64
	price_per_book      sql.NullFloat64
	sec_filings         sql.NullString
}

func GetStocks(c *gin.Context) {
	db, err := GetDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		panic(err)
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM snp_500_financials")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		panic(err)
	}
	defer results.Close()

	var stocks []STOCK
	for results.Next() {
		var stock STOCK

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

	if len(symbol) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid Stock Symbol"})
		return
	}

	db, err := GetDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		panic(err)
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM snp_500_financials WHERE symbol = ?", strings.ToUpper(symbol))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		panic(err)
	}
	defer results.Close()

	IsStockFound := false
	for results.Next() {
		IsStockFound = true
		var stock STOCK

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

		c.JSON(http.StatusOK, gin.H{
			"stock": stock,
		})
	}

	if !IsStockFound {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "Stock Not Found",
		})
	}
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
