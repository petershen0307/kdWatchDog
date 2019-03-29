package core

import "testing"

func Test_SaveKDValueToSheet_request_success(t *testing.T) {
	stockDailyInfo := []KDResult{
		{Date: 20150121, ClosePrice: 67.25, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
		{Date: 20150122, ClosePrice: 67.6, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
		{Date: 20150124, ClosePrice: 68.7, NHighPrice: 1.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
	}
	data := []KDStockInfo{
		KDStockInfo{
			StockID:      "1234",
			LatestKDInfo: stockDailyInfo[len(stockDailyInfo)-1],
			StockName:    "Just test",
		},
	}
	SaveKDValueToSheet(data, MonthPricePeriod)
}

func Test_GetStockWatchList(t *testing.T) {
	stockIDList := GetStockWatchList()
	t.Log(stockIDList)
}
