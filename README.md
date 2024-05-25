# stocks-server

**Run Server**

```bash
go run .
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

**Seed Database(with sample data)** _Alternatively you can use `db.sql` file_

```sql
INSERT INTO snp_500_financials (
  symbol,
  name,
  sector,
  price,
  price_per_earnings,
  dividend_yield,
  earnings_per_share,
  fifty_two_week_low,
  fifty_two_week_high,
  market_cap,
  EBITDA,
  price_per_sales,
  price_per_book,
  sec_filings
) VALUES
('AAPL', 'Apple Inc.', 'Technology', 145.09, 28.45, 0.007, 5.11, 107.32, 157.26, 2.41e12, 84.69e9, 7.23, 33.45, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0000320193'),
('MSFT', 'Microsoft Corp.', 'Technology', 258.74, 35.12, 0.009, 7.74, 198.14, 282.30, 1.93e12, 75.25e9, 10.12, 14.34, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0000789019'),
('GOOGL', 'Alphabet Inc.', 'Communication Services', 2378.00, 30.22, 0.0, 93.81, 2100.00, 2441.00, 1.59e12, 91.71e9, 8.95, 5.98, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0001652044'),
('AMZN', 'Amazon.com Inc.', 'Consumer Discretionary', 3342.88, 61.20, 0.0, 52.56, 2881.00, 3552.25, 1.70e12, 48.15e9, 4.21, 17.68, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0001018724'),
('FB', 'Facebook Inc.', 'Communication Services', 327.66, 25.43, 0.0, 10.87, 244.61, 378.51, 0.93e12, 32.67e9, 10.34, 6.48, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0001326801'),
('TSLA', 'Tesla Inc.', 'Consumer Discretionary', 680.76, 1125.40, 0.0, 0.64, 273.00, 900.40, 0.65e12, 7.22e9, 21.07, 37.22, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0001318605'),
('BRK.B', 'Berkshire Hathaway Inc.', 'Financials', 278.26, 25.32, 0.003, 11.00, 202.13, 293.43, 0.63e12, 40.82e9, 2.59, 1.37, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0001067983'),
('JNJ', 'Johnson & Johnson', 'Health Care', 164.30, 28.42, 0.025, 6.12, 133.65, 173.65, 0.43e12, 25.10e9, 5.13, 6.04, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0000200406'),
('V', 'Visa Inc.', 'Information Technology', 235.30, 43.30, 0.006, 5.43, 179.23, 250.93, 0.51e12, 18.90e9, 22.90, 11.23, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0001403161'),
('WMT', 'Walmart Inc.', 'Consumer Staples', 140.32, 29.87, 0.015, 4.70, 117.01, 153.66, 0.40e12, 25.57e9, 0.76, 4.42, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0000104169'),
('JPM', 'JPMorgan Chase & Co.', 'Financials', 151.23, 12.58, 0.027, 12.02, 91.38, 159.03, 0.48e12, 55.21e9, 3.94, 1.76, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0000019617'),
('MA', 'Mastercard Inc.', 'Information Technology', 372.15, 52.44, 0.004, 7.10, 280.20, 398.25, 0.37e12, 17.68e9, 19.88, 51.07, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0001141391'),
('PG', 'Procter & Gamble Co.', 'Consumer Staples', 137.61, 25.73, 0.023, 5.35, 121.54, 146.92, 0.33e12, 18.52e9, 4.34, 7.15, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0000080424'),
('NVDA', 'NVIDIA Corporation', 'Information Technology', 589.38, 91.23, 0.001, 6.46, 323.00, 648.57, 0.36e12, 8.27e9, 22.42, 21.87, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0001045810'),
('DIS', 'Walt Disney Co.', 'Communication Services', 179.12, 46.29, 0.0, 3.87, 115.14, 203.02, 0.32e12, 13.10e9, 4.85, 3.54, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0001744489'),
('HD', 'Home Depot Inc.', 'Consumer Discretionary', 318.73, 24.67, 0.021, 13.14, 246.59, 345.69, 0.34e12, 18.92e9, 2.44, 178.39, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0000354950'),
('VZ', 'Verizon Communications Inc.', 'Communication Services', 58.24, 13.24, 0.046, 4.40, 49.69, 61.95, 0.24e12, 48.07e9, 1.88, 4.58, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0000732717'),
('PYPL', 'PayPal Holdings Inc.', 'Information Technology', 261.57, 80.45, 0.0, 3.25, 160.00, 309.14, 0.31e12, 4.18e9, 19.50, 15.72, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0001633917'),
('ADBE', 'Adobe Inc.', 'Information Technology', 512.45, 48.72, 0.0, 10.51, 420.78, 536.88, 0.24e12, 7.45e9, 20.50, 13.57, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0000796343'),
('NFLX', 'Netflix Inc.', 'Communication Services', 503.55, 83.74, 0.0, 6.01, 384.00, 593.29, 0.23e12, 9.36e9, 8.74, 20.01, 'https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0001065280');
```

### cURL Commands

**Get all Stocks**

```bash
curl http://localhost:8080/stocks
```

**Get a Stock**

```bash
curl http://localhost:8080/stocks/NFLX
```

**Create Stock**

```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "symbol": "NFLX",
  "name": "Netflix Inc.",
  "sector": "Communication Services",
  "price": 503.55,
  "price_per_earnings": 83.74,
  "dividend_yield": 0.0,
  "earnings_per_share": 6.01,
  "fifty_two_week_low": 384.00,
  "fifty_two_week_high": 593.29,
  "market_cap": 0.23e12,
  "EBITDA": 9.36e9,
  "price_per_sales": 8.74,
  "price_per_book": 20.01,
  "sec_filings": "https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0001065280"
}' http://localhost:8080/stocks

```

**Update Stock**

```bash
curl -X PUT -H "Content-Type: application/json" -d '{
  "name": "Netflix Inc",
  "sector": "Communication Services",
  "price": 503.55,
  "price_per_earnings": 83.74,
  "dividend_yield": 0,
  "earnings_per_share": 6.01,
  "fifty_two_week_low": 384,
  "fifty_two_week_high": 593.29,
  "market_cap": 230000000000,
  "EBITDA": 9360000000,
  "price_per_sales": 8.74,
  "price_per_book": 20.01,
  "sec_filings": "https://www.sec.gov/cgi-bin/browse-edgar?action=getcompany&CIK=0001065280"
}' http://localhost:8080/stocks/NFLX
```

**Delete Stock**

```bash
curl -X DELETE http://localhost:8080/stocks/NFLX
```
