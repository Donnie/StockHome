# Stockhome
Go app to find and store daily stock data.

Dockerised Go app which updates candle data with daily resolution.

This app tracks the S&P 500 Index and updates the stocks list weekly. And it updates the candle data for all stocks in DB daily at 01:00 UTC.

Data is locally stored in sqlite file in db/sql.db

## Future
- API with secret key
- Telegram Bot to check stock historical prices
