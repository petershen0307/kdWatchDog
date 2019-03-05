package core

import "testing"

func Test_saveDataToGoogleSheet_request_success(t *testing.T) {
	stockDailyInfo := []KDResult{
		{Date: 20150121, ClosePrice: 67.25, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
		{Date: 20150122, ClosePrice: 67.6, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
		{Date: 20150123, ClosePrice: 68.7, NHighPrice: 0.0, NLowPrice: 0, RSV: 0.0, K: 50, D: 50},
	}
	SaveDataToGoogleSheet(stockDailyInfo, "i am test", DailyPricePeriod)
}
