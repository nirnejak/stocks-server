package main

type STOCK struct {
	Symbol              string  `json:"symbol"`
	Name                string  `json:"name"`
	Sector              string  `json:"sector"`
	Price               float64 `json:"price"`
	PricePerEarnings    float64 `json:"price_per_earnings"`
	DividendYield       float64 `json:"dividend_yield"`
	EarningsPerShare    float64 `json:"earnings_per_share"`
	FiftyTwoWeekLow     float64 `json:"fifty_two_week_low"`
	FiftyTwoWeekHigh    float64 `json:"fifty_two_week_high"`
	MarketCap           float64 `json:"market_cap"`
	EBITDA              float64 `json:"EBITDA"`
	PricePerSales       float64 `json:"price_per_sales"`
	PricePerBook        float64 `json:"price_per_book"`
	SecFilings          string  `json:"sec_filings"`
}
