package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// GetHTTP to make GET call
func GetHTTP(api string, ratePM *float64) string {
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

// StringToCSV transforms csv string to array
func StringToCSV(input string) (output [][]string) {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	for _, line := range lines {
		output = append(output, strings.Split(line, ","))
	}
	return
}

// RoundTimeToDate round times to day start
func RoundTimeToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}
