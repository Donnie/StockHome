package api

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Candle struct to hold OHCLV data
type Candle struct {
	Closes     []float64 `json:"c"`
	Highs      []float64 `json:"h"`
	Lows       []float64 `json:"l"`
	Opens      []float64 `json:"o"`
	Status     string    `json:"s"`
	Timestamps []int64   `json:"t"`
	Volumes    []float64 `json:"v"`
}

// GetCandleLatest generates the REST API for the data provider
func GetCandleLatest(symbol string, lastTime time.Time, free bool) (string, *float64) {
	api := os.Getenv("API")
	apikey := os.Getenv("KEY")
	endUnix := RoundTimeToDate(time.Now()).Unix()
	lastUnix := RoundTimeToDate(lastTime).Unix()
	ep := fmt.Sprintf(
		"%s?symbol=%s&resolution=D&from=%d&to=%d&token=%s",
		api,
		symbol,
		lastUnix,
		endUnix,
		apikey,
	)
	if free {
		rate := 59.0
		return ep, &rate
	}
	return ep, nil
}

// ReadToStruct transforms candle json to struct
func ReadToStruct(symbol string, lastTime *time.Time, free bool) (res Candle) {
	if lastTime == nil {
		// get one year prev
		tyme := time.Now().AddDate(-1, -1, 0)
		lastTime = &tyme
	}
	input := GetHTTP(GetCandleLatest(symbol, *lastTime, free))
	err := json.Unmarshal([]byte(input), &res)
	if err != nil {
		fmt.Printf("Failed to read res, err: %v", err)
	}
	return
}
