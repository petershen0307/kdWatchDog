package core

import "testing"

func Test_saveKDValueToSheet_request_success(t *testing.T) {
	stockDailyInfo := []KDResult{
		{Date: 20150121, ClosePrice: 67.25, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
		{Date: 20150122, ClosePrice: 67.6, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
		{Date: 20150124, ClosePrice: 68.7, NHighPrice: 1.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
	}
	data := []KDStockInfo{
		KDStockInfo{
			stockID:      "1234",
			latestKDInfo: stockDailyInfo[len(stockDailyInfo)-1],
		},
	}
	saveKDValueToSheet(data, MonthPricePeriod)
}
