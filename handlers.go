package main

import (
	"os"

	"github.com/Donnie/stockhome/models"
	"github.com/gin-gonic/gin"
)

func stockHandler(c *gin.Context) {
	key, ok := c.GetQuery("key")
	start, _ := c.GetQuery("start")
	end, _ := c.GetQuery("end")

	if ok && key == os.Getenv("PASS") {
		stock := models.GetStockWithCandles(c.Param("sym"), start, end)
		c.JSON(200, stock)
		return
	}
	c.JSON(403, "Please provide API key")
}

func indexHandler(c *gin.Context) {
	key, ok := c.GetQuery("key")

	if ok && key == os.Getenv("PASS") {
		index := models.GetIndexWithStocks(c.Param("sym"))
		c.JSON(200, index)
		return
	}
	c.JSON(403, "Please provide API key")
}
