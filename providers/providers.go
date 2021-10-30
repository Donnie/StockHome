package providers

import (
	"fmt"
	"math"
	"time"

	"github.com/Donnie/stockhome/api"
	"github.com/Donnie/stockhome/models"
	"github.com/Donnie/stockhome/ptr"
	"gorm.io/gorm"
)

// UpdateIndices updates list of stocks from an Index
func UpdateIndices(db *gorm.DB) {
	fmt.Println("Indices update started:", time.Now())

	UpdateIdxS500(db)
}

// UpdateCandles updates stock data
func UpdateCandles(db *gorm.DB) {
	fmt.Println("Candles update started:", time.Now())

	// get stocks
	stocks := []models.Stock{}
	db.Find(&stocks)

	// for each stock
	for _, stock := range stocks {
		// get last candle of stock
		lastCandle := models.Candle{}
		db.Where("stock_id = ?", stock.ID).Order("date DESC").First(&lastCandle)

		if lastCandle.Date != nil &&
			lastCandle.Date.Unix() == RoundTimeToDate(time.Now()).Unix() {
			continue
		}

		// fetch data from provider
		candles := api.ReadToStruct(*stock.Symbol, lastCandle.Date, true)
		newCandles := []models.Candle{}

		// process data to models
		for i, date := range candles.Timestamps {
			candle := models.Candle{
				Close:   ptr.Int64(int64(math.Round(candles.Closes[i] * 100.0))),
				Date:    ptr.Time(time.Unix(date, 0)),
				High:    ptr.Int64(int64(math.Round(candles.Highs[i] * 100.0))),
				Low:     ptr.Int64(int64(math.Round(candles.Lows[i] * 100.0))),
				Open:    ptr.Int64(int64(math.Round(candles.Opens[i] * 100.0))),
				StockID: &stock.ID,
				Volume:  ptr.Int64(int64(math.Round(candles.Volumes[i]))),
			}
			if lastCandle.Date != nil &&
				lastCandle.Date.Unix() == candle.Date.Unix() {
				// skip if dates are same
				continue
			}
			newCandles = append(newCandles, candle)
		}

		if len(newCandles) > 0 {
			// upload data to DB
			db.Create(&newCandles)
		}
	}
}

// UpdateHistoryCandles updates stock data
func UpdateHistoryCandles(db *gorm.DB) {
	// Find one stock with missing candles before a certain date
	stock := models.Stock{}

	db.Where(`history = 0`).First(&stock)

	if stock.Symbol == nil {
		return
	}

	fmt.Println("Stock update started:", *stock.Symbol, time.Now())

	// get first candle of stock
	firstCandle := models.Candle{}
	db.Where("stock_id = ?", stock.ID).Order("date ASC").First(&firstCandle)

	// fetch data from provider
	candles := api.GetHistory(*stock.Symbol, stock.ID, firstCandle.Date)

	if len(candles) > 0 {
		// upload data to DB
		db.Create(&candles)
	}

	db.Model(&stock).Update("history", 1)
}
