package core

import (
	"fmt"
	"log"

	sheets "google.golang.org/api/sheets/v4"
)

const spreadsheetID = "1VWsWERMQWhycfxz7xZEK0I8rBSHOwnW09OCG7HIL1k4"

// KDStockInfo include stock id and kd info
type KDStockInfo struct {
	StockID      string
	LatestKDInfo KDResult
	StockName    string
}

//SaveKDValueToSheet send KD info to google sheet
func SaveKDValueToSheet(stockData []KDStockInfo, sheetName pricePeriod) {
	//日期	股票	名稱	收盤價	最高價	最低價	RSV	K	D
	//A     B       C      D      E       F      G   H   I
	sheetData := sheets.ValueRange{
		Range:          fmt.Sprintf("%v!A2:I", sheetName),
		MajorDimension: "ROWS",
	}
	for _, oneStock := range stockData {
		rowData := []interface{}{
			oneStock.LatestKDInfo.Date,
			oneStock.StockID,
			oneStock.StockName,
			oneStock.LatestKDInfo.ClosePrice,
			oneStock.LatestKDInfo.NHighPrice,
			oneStock.LatestKDInfo.NLowPrice,
			oneStock.LatestKDInfo.RSV,
			oneStock.LatestKDInfo.K,
			oneStock.LatestKDInfo.D,
		}
		sheetData.Values = append(sheetData.Values, rowData)
	}
	data := []*sheets.ValueRange{
		&sheetData,
	}
	rb := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "RAW",
		Data:             data,
		ResponseValueRenderOption: "UNFORMATTED_VALUE",
		//IncludeValuesInResponse:   false,
	}
	sheetSvc := getSheetSvc()
	_, err := sheetSvc.Spreadsheets.Values.BatchUpdate(spreadsheetID, rb).Do()
	if err != nil {
		log.Fatal(err)
	}
}
