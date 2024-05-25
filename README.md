# stocks-server

**Run Server**

```bash
go run main.go
```

**Create a database with the script below**

```sql
CREATE TABLE snp_500_financials (
  symbol              TEXT,
	name                TEXT,
	sector              TEXT,
	price               DOUBLE PRECISION,
	price_per_earnings  DOUBLE PRECISION,
	dividend_yield      DOUBLE PRECISION,
	earnings_per_share  DOUBLE PRECISION,
	fifty_two_week_low  DOUBLE PRECISION,
	fifty_two_week_high DOUBLE PRECISION,
	market_cap          DOUBLE PRECISION,
	EBITDA              DOUBLE PRECISION,
	price_per_sales     DOUBLE PRECISION,
	price_per_book      DOUBLE PRECISION,
	sec_filings         TEXT
)

```
