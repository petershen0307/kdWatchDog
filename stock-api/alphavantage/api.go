package alphavantage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/petershen0307/kdWatchDog/models"
	stockapi "github.com/petershen0307/kdWatchDog/stock-api"
)

const baseURL = "https://www.alphavantage.co/query?"

// API integrate alphavantage API
type API struct {
	key         string
	intervalMap map[stockapi.ResolutionInterval]string
}

// New is the function to create alphavantage API object
func New(key string) *API {
	return &API{
		key: key,
		intervalMap: map[stockapi.ResolutionInterval]string{
			// interval: 1min, 5min, 15min, 30min, 60min, daily, weekly, monthly
			stockapi.OneMin:       "1min",
			stockapi.FiveMin:      "5min",
			stockapi.FifthteenMin: "15min",
			stockapi.ThirtyMin:    "30min",
			stockapi.SixtyMin:     "60min",
			stockapi.Daily:        "daily",
			stockapi.Weekly:       "weekly",
			stockapi.Monthly:      "monthly",
		},
	}
}

// GetStockSymbol get stock market symbol
func (a *API) GetStockSymbol(exchange string) []string {
	return []string{}
}

// GetSTOCH get kd indicator
func (a *API) GetSTOCH(symbol string, interval stockapi.ResolutionInterval, fastkperiod, slowkperiod, slowdperiod, slowkmatype, slowdmatype uint8) models.STOCH {
	result := models.STOCH{}
	//https://www.alphavantage.co/query?function=STOCH&symbol=AAPL&interval=daily&apikey=...
	apiArgs := url.Values{
		"function":    []string{"STOCH"},
		"symbol":      []string{strings.ToUpper(symbol)},
		"interval":    []string{a.intervalMap[interval]},
		"fastkperiod": []string{strconv.Itoa(int(fastkperiod))},
		"slowkperiod": []string{strconv.Itoa(int(slowkperiod))},
		"slowdperiod": []string{strconv.Itoa(int(slowdperiod))},
		"slowkmatype": []string{strconv.Itoa(int(slowkmatype))},
		"slowdmatype": []string{strconv.Itoa(int(slowdmatype))},
		"datatype":    []string{"json"},
		"apikey":      []string{a.key},
	}
	requestURL := baseURL + apiArgs.Encode()
	log.Printf("alphavantage request url=%v", requestURL)
	if resp, err := http.Get(requestURL); err == nil {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Printf("alphavantage http status=%v, msg=%v", resp.StatusCode, resp.Status)
			return result
		}
		data, _ := ioutil.ReadAll(resp.Body)
		jsonMap := map[string]interface{}{}
		json.Unmarshal(data, &jsonMap)
		lastDate := ""
		if jMap, ok := jsonMap["Meta Data"].(map[string]interface{}); ok {
			lastDate = jMap["3: Last Refreshed"].(string)
		}
		if jMap, ok := jsonMap["Technical Analysis: STOCH"].(map[string]interface{}); ok {
			result.K = jMap[lastDate].(map[string]interface{})["SlowK"].(string)
			result.D = jMap[lastDate].(map[string]interface{})["SlowD"].(string)
		}
	} else {
		log.Printf("alphavantage http connect error=%v", err)
	}
	return result
}

func (a *API) getPrice(symbol, priceFunction string, decoder func(map[string]interface{}) models.Price) models.Price {
	result := models.Price{}
	//https://www.alphavantage.co/query?function=TIME_SERIES_MONTHLY&symbol=IBM&apikey=demo
	apiArgs := url.Values{
		"function": []string{priceFunction},
		"symbol":   []string{strings.ToUpper(symbol)},
		"datatype": []string{"json"},
		"apikey":   []string{a.key},
	}
	requestURL := baseURL + apiArgs.Encode()
	log.Printf("alphavantage request url=%v", requestURL)
	if resp, err := http.Get(requestURL); err == nil {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Printf("alphavantage http status=%v, msg=%v", resp.StatusCode, resp.Status)
			return result
		}
		data, _ := ioutil.ReadAll(resp.Body)
		jsonMap := map[string]interface{}{}
		json.Unmarshal(data, &jsonMap)
		result = decoder(jsonMap)
	} else {
		log.Printf("alphavantage http connect error=%v", err)
	}
	return result
}

// GetDailyPrice daily price
func (a *API) GetDailyPrice(symbol string) models.Price {
	r := a.getPrice(symbol, "TIME_SERIES_DAILY", func(jsonMap map[string]interface{}) models.Price {
		lastDate := ""
		if jMap, ok := jsonMap["Meta Data"].(map[string]interface{}); ok {
			lastDate = jMap["3. Last Refreshed"].(string)
		}
		if jMap, ok := jsonMap["Time Series (Daily)"].(map[string]interface{}); ok {
			return models.Price{
				Open:  jMap[lastDate].(map[string]interface{})["1. open"].(string),
				High:  jMap[lastDate].(map[string]interface{})["2. high"].(string),
				Low:   jMap[lastDate].(map[string]interface{})["3. low"].(string),
				Close: jMap[lastDate].(map[string]interface{})["4. close"].(string),
			}
		}
		return models.Price{}
	})
	return r
}

// GetWeeklyPrice weekly price
func (a *API) GetWeeklyPrice(symbol string) models.Price {
	r := a.getPrice(symbol, "TIME_SERIES_WEEKLY", func(jsonMap map[string]interface{}) models.Price {
		lastDate := ""
		if jMap, ok := jsonMap["Meta Data"].(map[string]interface{}); ok {
			lastDate = jMap["3. Last Refreshed"].(string)
		}
		if jMap, ok := jsonMap["Weekly Time Series"].(map[string]interface{}); ok {
			return models.Price{
				Open:  jMap[lastDate].(map[string]interface{})["1. open"].(string),
				High:  jMap[lastDate].(map[string]interface{})["2. high"].(string),
				Low:   jMap[lastDate].(map[string]interface{})["3. low"].(string),
				Close: jMap[lastDate].(map[string]interface{})["4. close"].(string),
			}
		}
		return models.Price{}
	})
	return r
}

// GetMonthlyPrice monthly price
func (a *API) GetMonthlyPrice(symbol string) models.Price {
	r := a.getPrice(symbol, "TIME_SERIES_MONTHLY", func(jsonMap map[string]interface{}) models.Price {
		lastDate := ""
		if jMap, ok := jsonMap["Meta Data"].(map[string]interface{}); ok {
			lastDate = jMap["3. Last Refreshed"].(string)
		}
		if jMap, ok := jsonMap["Monthly Time Series"].(map[string]interface{}); ok {
			return models.Price{
				Open:  jMap[lastDate].(map[string]interface{})["1. open"].(string),
				High:  jMap[lastDate].(map[string]interface{})["2. high"].(string),
				Low:   jMap[lastDate].(map[string]interface{})["3. low"].(string),
				Close: jMap[lastDate].(map[string]interface{})["4. close"].(string),
			}
		}
		return models.Price{}
	})
	return r
}
