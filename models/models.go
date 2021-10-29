package models

import (
	"database/sql"
	"time"
)

// Candle to hold daily OHCLV data
type Candle struct {
	ID        uint         `json:"-" gorm:"primarykey"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`

	Close   *int64     `json:"close"`
	Date    *time.Time `json:"date"`
	High    *int64     `json:"high"`
	Low     *int64     `json:"low"`
	Open    *int64     `json:"open"`
	StockID *uint      `json:"-"`
	Volume  *int64     `json:"volume"`
}

// Stock to hold stock Info
type Stock struct {
	ID        uint         `json:"-" gorm:"primarykey"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`

	Candles     []Candle `json:"candles"`
	Description *string  `json:"description"`
	Name        *string  `json:"name"`
	Sector      *string  `json:"sector"`
	Symbol      *string  `json:"symbol"`
}

// Index to hold Indices
type Index struct {
	ID        uint         `json:"-" gorm:"primarykey"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
	DeletedAt sql.NullTime `json:"-" gorm:"index"`

	Name        *string `json:"name"`
	Description *string `json:"description"`
	Stocks      []Stock `json:"stocks" gorm:"many2many:indices_stocks"`
	Symbol      *string `json:"symbol"`
}

// TableName overrides the table name used by Index to Indices
func (Index) TableName() string {
	return "indices"
}
