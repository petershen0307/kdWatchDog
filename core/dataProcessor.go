package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

type pricePeriod string

const (
	monthPricePeriod pricePeriod = "m"
	dailyPricePeriod pricePeriod = "d"
	weekPricePeriod  pricePeriod = "w"
)

const stockPriceURLTemplate string = "https://tw.quote.finance.yahoo.net/quote/q?type=ta&perd=%v&mkt=10&sym=%v"

// GetStockInfoFromWeb will return StockPriceInfo which download from yahoo stock web
func GetStockInfoFromWeb(stockID uint, period pricePeriod) (StockPriceInfo, error) {
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

func downloadStockPrice(stockID uint, period pricePeriod) (string, error) {
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
	OpenPrice  float32 `json:"o"`
	HighPrice  float32 `json:"h"`
	LowPrice   float32 `json:"l"`
	ClosePrice float32 `json:"c"`
	Volume     uint    `json:"v"`
}
type stockMemo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// StockPriceInfo contain stock price information which are daily or monthly
type StockPriceInfo struct {
	ID        string      `json:"id"`
	Period    pricePeriod `json:"perd"`
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
