package models

import "gorm.io/gorm"

// GetStockWithCandles returns stock data with candles
func GetStockWithCandles(sym string, start, end string) (stock Stock) {
	DB.Preload("Candles", func(db *gorm.DB) *gorm.DB {
		if start != "" {
			db = db.Where("date > ?", start)
		}
		if end != "" {
			db = db.Where("date < ?", end)
		}
		return db.Order("date ASC")
	}).
		Where("symbol = ?", sym).
		Find(&stock)
	return
}

// GetIndexWithStocks returns index data with stocks
func GetIndexWithStocks(sym string) (index Index) {
	DB.Preload("Stocks", func(db *gorm.DB) *gorm.DB {
		return db.Order("symbol ASC")
	}).
		Where("symbol = ?", sym).
		Find(&index)
	return
}
