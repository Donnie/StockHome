# StockHome
I wrote this app to collect and store historical OHCLV candle data with daily resolution.

The app regularly updates the stocks data by 21:30 UTC.

As of now the focus is on S&P 500 stocks, but there are plans to extend it to other indices.

Candle data is available on /stocks/:sym endpoint

Index data is available on /indices/:sym endpoint

A *special key* has to be added as query parameter, which can be set in PASS variable in .env file

## Data
You can get access to the stock data with daily resolution in USD terms, split adjusted, going back to 2010-01-04.

API: https://stockhome.donnie.in

API example: https://stockhome.donnie.in/stocks/MMM?key=admin&start=2020-01-01&end=2020-12-31
- start and end are optional parameters

Please request access key by emailing to stockhome@donnie.in with subject: OHCLV Data.

Please do not abuse.

## Future
- Data going back to 2000-01-01
- Telegram Bot to check stock historical prices
- Add FTSE 100 Data
