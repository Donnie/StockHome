package models

import "gorm.io/gorm"

// GetStockWithCandles returns stock data with candles
func GetStockWithCandles(db *gorm.DB, sym string) (stock Stock) {
	db.Preload("Candles", func(db *gorm.DB) *gorm.DB {
		return db.Order("date ASC")
	}).
		Where("symbol = ?", sym).
		Find(&stock)
	return
}
