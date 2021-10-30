package api

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/Donnie/stockhome/models"
	"github.com/Donnie/stockhome/ptr"
)

// GenAlpha generates the REST API
func GenAlpha(symbol string, free bool) (string, *float64) {
	api := os.Getenv("HIST")
	apikey := os.Getenv("HIST_KEY")
	ep := fmt.Sprintf(
		"%s?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s&datatype=csv&outputsize=full",
		api,
		symbol,
		apikey,
	)
	if free {
		rate, _ := strconv.ParseFloat(os.Getenv("HIST_RATE"), 64)
		return ep, &rate
	}
	return ep, nil
}

// GetHistory sorts
func GetHistory(sym string, stockID uint, firstDate *time.Time) (candles []models.Candle) {
	data := GetHTTP(GenAlpha(sym, false))
	csv := StringToCSV(data)

	for _, cs := range csv[1:] {
		if len(cs) < 6 {
			continue
		}
		date, _ := time.Parse("2006-01-02", cs[0])
		start, _ := time.Parse("2006-01-02", os.Getenv("HIST_START"))
		if date.Unix() < start.Unix() {
			continue
		}
		if firstDate != nil && date.Unix() >= firstDate.Unix() {
			continue
		}
		open, err := strconv.ParseFloat(cs[1], 64)
		if err != nil {
			continue
		}
		high, err := strconv.ParseFloat(cs[2], 64)
		if err != nil {
			continue
		}
		low, err := strconv.ParseFloat(cs[3], 64)
		if err != nil {
			continue
		}
		close, err := strconv.ParseFloat(cs[4], 64)
		if err != nil {
			continue
		}
		volume, err := strconv.ParseFloat(cs[5], 64)
		if err != nil {
			continue
		}
		candle := models.Candle{
			Date:    ptr.Time(date),
			Open:    ptr.Int64(int64(math.Round(open * 100.0))),
			High:    ptr.Int64(int64(math.Round(high * 100.0))),
			Low:     ptr.Int64(int64(math.Round(low * 100.0))),
			Close:   ptr.Int64(int64(math.Round(close * 100.0))),
			Volume:  ptr.Int64(int64(math.Round(volume))),
			StockID: &stockID,
		}
		candles = append(candles, candle)
	}
	return
}
