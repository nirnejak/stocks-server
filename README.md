# stocks-server

**Run Server**

```bash
go run main.go
```

**Create a database with the script below**

```sql
CREATE TABLE snp_500_financials (
  symbol              VARCHAR(255),
  name                VARCHAR(255),
  sector              VARCHAR(255),
  price               DOUBLE,
  price_per_earnings  DOUBLE,
  dividend_yield      DOUBLE,
  earnings_per_share  DOUBLE,
  fifty_two_week_low  DOUBLE,
  fifty_two_week_high DOUBLE,
  market_cap          DOUBLE,
  EBITDA              DOUBLE,
  price_per_sales     DOUBLE,
  price_per_book      DOUBLE,
  sec_filings         VARCHAR(255)
);

```
