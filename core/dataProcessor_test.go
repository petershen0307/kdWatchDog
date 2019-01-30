package core

import (
	"reflect"
	"testing"
)

func Test_parseStockPriceJSON_WhenCall_Success(t *testing.T) {
	gotStr := `null({"mkt":"10","id":"1234","perd":"d","type":"ta","mem":{"id":"1234","name":"黑松","125":29.95,"126":29.9,"638":0.0,"127":0.0},"ta":[{"t":20180928,"o":31.0,"h":31.0,"l":30.75,"c":30.8,"v":48},{"t":20181001,"o":30.85,"h":30.9,"l":30.75,"c":30.9,"v":78}]});`
	expectedResult := StockPriceInfo{
		ID:     "1234",
		Period: dailyPricePeriod,
		Mem: stockMemo{
			ID:   "1234",
			Name: "黑松",
		},
		PriceInfo: []dailyInfo{
			{
				Date:       20180928,
				OpenPrice:  31.0,
				HighPrice:  31.0,
				LowPrice:   30.75,
				ClosePrice: 30.8,
				Volume:     48,
			},
			{
				Date:       20181001,
				OpenPrice:  30.85,
				HighPrice:  30.9,
				LowPrice:   30.75,
				ClosePrice: 30.9,
				Volume:     78,
			},
		},
	}
	stockDailyInfo, err := parseStockPriceJSON(gotStr)
	if err != nil {
		t.Errorf("Result was incorrect, got %v", err)
	}
	if !reflect.DeepEqual(expectedResult, stockDailyInfo) {
		t.Errorf("Result was incorrect\ngot: %v\nwant:%v", stockDailyInfo, expectedResult)
	}
}
