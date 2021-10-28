package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// GetHTTP to make GET call
func GetHTTP(api string, ratePM *int) string {
	fmt.Println(api)
	resp, err := http.Get(api)
	if err != nil {
		fmt.Printf("API Endpoint not reachable, err: %v", err)
		os.Exit(3)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("Access error, status: %d", resp.StatusCode)
		os.Exit(3)
	}

	if ratePM != nil {
		// respect rate limiting
		rate := time.Duration(60000 / *ratePM)
		time.Sleep(rate * time.Millisecond)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Body not readable, err: %v", err)
		os.Exit(3)
	}
	return string(body)
}

// RoundTimeToDate round times to day start
func RoundTimeToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}
