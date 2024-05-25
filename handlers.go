package main

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetStocks(c *gin.Context) {
	db := GetDB()
	rows, err := db.Query("SELECT * FROM snp_500_financials")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var stocks []STOCK
	for rows.Next() {
		var stock STOCK
		if err := rows.Scan(
			&stock.Symbol,
			&stock.Name,
			&stock.Sector,
			&stock.Price,
			&stock.PricePerEarnings,
			&stock.DividendYield,
			&stock.EarningsPerShare,
			&stock.FiftyTwoWeekLow,
			&stock.FiftyTwoWeekHigh,
			&stock.MarketCap,
			&stock.EBITDA,
			&stock.PricePerSales,
			&stock.PricePerBook,
			&stock.SecFilings,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		stocks = append(stocks, stock)
	}
	c.JSON(http.StatusOK, gin.H{"stocks": stocks})
}

func GetStock(c *gin.Context) {
	symbol := c.Param("symbol")

	if len(symbol) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Stock Symbol"})
		return
	}

	db := GetDB()
	row := db.QueryRow("SELECT * FROM snp_500_financials WHERE symbol = ?", strings.ToUpper(symbol))

	var stock STOCK
	if err := row.Scan(
		&stock.Symbol,
		&stock.Name,
		&stock.Sector,
		&stock.Price,
		&stock.PricePerEarnings,
		&stock.DividendYield,
		&stock.EarningsPerShare,
		&stock.FiftyTwoWeekLow,
		&stock.FiftyTwoWeekHigh,
		&stock.MarketCap,
		&stock.EBITDA,
		&stock.PricePerSales,
		&stock.PricePerBook,
		&stock.SecFilings,
	); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Stock Not Found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"stock": stock})
}

func CreateStock(c *gin.Context) {
	var stock STOCK
	if err := c.BindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	db := GetDB()
	result, err := db.Exec("INSERT INTO snp_500_financials (symbol, name, sector, price, price_per_earnings, dividend_yield, earnings_per_share, fifty_two_week_low, fifty_two_week_high, market_cap, EBITDA, price_per_sales, price_per_book, sec_filings) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		stock.Symbol,
		stock.Name,
		stock.Sector,
		stock.Price,
		stock.PricePerEarnings,
		stock.DividendYield,
		stock.EarningsPerShare,
		stock.FiftyTwoWeekLow,
		stock.FiftyTwoWeekHigh,
		stock.MarketCap,
		stock.EBITDA,
		stock.PricePerSales,
		stock.PricePerBook,
		stock.SecFilings,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stock"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Stock created successfully"})
}

func UpdateStock(c *gin.Context) {
	symbol := c.Param("symbol")

	// Implementation for updating a stock in the database

	c.JSON(http.StatusOK, gin.H{
		"symbol":  symbol,
		"message": "Stock Updated",
	})
}

func DeleteStock(c *gin.Context) {
	symbol := c.Param("symbol")

	db := GetDB()
	result, err := db.Exec("DELETE FROM snp_500_financials WHERE symbol = ?", symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"symbol":  symbol,
		"message": "Stock Deleted",
	})
}
