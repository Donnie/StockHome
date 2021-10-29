package providers

import (
	"fmt"
	"os"

	"github.com/Donnie/stockhome/models"
	"github.com/Donnie/stockhome/ptr"
	"gorm.io/gorm"
)

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
