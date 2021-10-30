package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Donnie/stockhome/models"
	"github.com/Donnie/stockhome/providers"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
)

func init() {
	if _, err := os.Stat(".env.local"); err == nil {
		godotenv.Load(".env.local")
		fmt.Println("Running for " + os.Getenv("ENV"))
		return
	}
	if _, err := os.Stat(".env"); err == nil {
		godotenv.Load(".env")
		fmt.Println("Running for " + os.Getenv("ENV"))
		return
	}
	fmt.Println(".env file not found")
	os.Exit(3)
}

func main() {
	// put DB to context
	db := initDB()

	r := gin.Default()
	r.GET("/stocks/:sym", func(c *gin.Context) {
		key, ok := c.GetQuery("key")

		if ok && key == os.Getenv("PASS") {
			stock := models.GetStockWithCandles(db, c.Param("sym"))
			c.JSON(200, stock)
			return
		}
		c.JSON(403, "Please provide API key")
	})

	r.GET("/indices/:sym", func(c *gin.Context) {
		key, ok := c.GetQuery("key")

		if ok && key == os.Getenv("PASS") {
			index := models.GetIndexWithStocks(db, c.Param("sym"))
			c.JSON(200, index)
			return
		}
		c.JSON(403, "Please provide API key")
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, nil)
	})
	go r.Run()

	// create scheduler
	// update indices
	gocron.Every(1).Week().Do(providers.UpdateIndices, db)

	// update stocks
	gocron.Every(1).Day().At(os.Getenv("TIME")).Do(providers.UpdateCandles, db)

	// Update historical candles
	rate, _ := strconv.Atoi(os.Getenv("HIST_RATE"))
	gocron.Every(uint64(rate)).Second().Do(providers.UpdateHistoryCandles, db)

	// Start jobs
	<-gocron.Start()
}
