package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

type PricePeriod string

const (
	// MonthPricePeriod is month period
	MonthPricePeriod PricePeriod = "m"
	// DailyPricePeriod is daily period
	DailyPricePeriod PricePeriod = "d"
	// WeekPricePeriod is week period
	WeekPricePeriod PricePeriod = "w"
)

const stockPriceURLTemplate string = "https://tw.quote.finance.yahoo.net/quote/q?type=ta&perd=%v&mkt=10&sym=%v"

// GetStockInfoFromWeb will return StockPriceInfo which download from yahoo stock web
func GetStockInfoFromWeb(stockID string, period PricePeriod) (StockPriceInfo, error) {
	rawData, err := downloadStockPrice(stockID, period)
	if err != nil {
		return StockPriceInfo{}, err
	}
	stockInfo, err := parseStockPriceJSON(rawData)
	if err != nil {
		return StockPriceInfo{}, err
	}
	return stockInfo, nil
}

func downloadStockPrice(stockID string, period PricePeriod) (string, error) {
	timeoutRequest := http.Client{Timeout: time.Minute * 5}
	url := fmt.Sprintf(stockPriceURLTemplate, period, stockID)
	response, err := timeoutRequest.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if http.StatusOK != response.StatusCode {
		return "", fmt.Errorf("Got http error code %v", response.StatusCode)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	return buf.String(), nil
}

type dailyInfo struct {
	Date       uint    `json:"t"`
	OpenPrice  float64 `json:"o"`
	HighPrice  float64 `json:"h"`
	LowPrice   float64 `json:"l"`
	ClosePrice float64 `json:"c"`
	Volume     uint    `json:"v"`
}
type stockMemo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// StockPriceInfo contain stock price information which are daily or monthly
type StockPriceInfo struct {
	ID        string      `json:"id"`
	Period    PricePeriod `json:"perd"`
	Mem       stockMemo   `json:"mem"`
	PriceInfo []dailyInfo `json:"ta"`
}

func parseStockPriceJSON(rawData string) (StockPriceInfo, error) {
	re := regexp.MustCompile("{.*}")
	jsonStr := re.FindString(rawData)
	var jsonData StockPriceInfo
	err := json.Unmarshal([]byte(jsonStr), &jsonData)
	if err != nil {
		return StockPriceInfo{}, err
	}
	return jsonData, nil
}
