package main

import (
	"fmt"
	"os"

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
	initDB()

	r := gin.Default()

	r.GET("/stocks/:sym", stockHandler)
	r.GET("/indices/:sym", indexHandler)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, nil)
	})

	go r.Run()

	// create scheduler
	// update indices
	gocron.Every(1).Week().Do(providers.UpdateIndices, models.DB)

	// update stocks
	gocron.Every(1).Day().At(os.Getenv("TIME")).Do(providers.UpdateCandles, models.DB)

	// Update historical candles
	// rate, _ := strconv.Atoi(os.Getenv("HIST_RATE"))
	// gocron.Every(uint64(rate)).Second().Do(providers.UpdateHistoryCandles, models.DB)

	// Start jobs
	<-gocron.Start()
}
