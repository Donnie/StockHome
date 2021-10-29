# Stockhome
Go app to find and store daily stock data.

Dockerised Go app which updates candle data with daily resolution.

This app tracks the S&P 500 Index and updates the stocks list weekly. And it updates the candle data for all stocks in DB daily at 01:00 UTC.

Data is locally stored in sqlite file in db/sql.db

## API
Candle data is available on /stocks/:sym endpoint

A special key has to be added as query parameter, which can be set in PASS variable in .env file

## Future
- Telegram Bot to check stock historical prices
