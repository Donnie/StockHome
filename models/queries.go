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

// GetIndexWithStocks returns index data with stocks
func GetIndexWithStocks(db *gorm.DB, sym string) (index Index) {
	db.Preload("Stocks", func(db *gorm.DB) *gorm.DB {
		return db.Order("symbol ASC")
	}).
		Where("symbol = ?", sym).
		Find(&index)
	return
}
