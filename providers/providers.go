package providers

import (
	"fmt"
	"math"
	"os"
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

// UpdateIdxS500 updates the S&P 500 index
func UpdateIdxS500(db *gorm.DB) {
	data := GetHTTP(os.Getenv("IDX_S500"), nil)

	// get indices
	index := models.Index{
		Description: ptr.String("The S&P 500 is a stock market index that is viewed as a measure of how well the stock market is performing overall. It includes around 500 of the largest U.S. companies."),
		Name:        ptr.String("S&P 500"),
		Symbol:      ptr.String("S500"),
	}

	// if S500 does not exist create it
	db.FirstOrCreate(&index, models.Index{Symbol: ptr.String("S500")})

	// Clear existing relations
	db.Model(&index).Association("Stocks").Clear()

	// Go one by one from Index data
	csv := StringToCSV(data)
	newStocks := []models.Stock{}
	for _, cs := range csv[1:] {
		if len(cs) > 2 {
			stock := models.Stock{
				Symbol: ptr.String(cs[0]),
				Name:   ptr.String(cs[1]),
				Sector: ptr.String(cs[2]),
			}

			// create the stock if it does not exist
			db.FirstOrCreate(&stock, models.Stock{Symbol: ptr.String(cs[0])})
			newStocks = append(newStocks, stock)

			// slow DB connection
			fmt.Println("Updated", cs[0])
		}
	}

	// build new relations
	db.Model(&index).Association("Stocks").Append(newStocks)
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
