package core

import (
	"fmt"
	"log"

	sheets "google.golang.org/api/sheets/v4"
)

const spreadsheetID = "1VWsWERMQWhycfxz7xZEK0I8rBSHOwnW09OCG7HIL1k4"

// KDStockInfo include stock id and kd info
type KDStockInfo struct {
	stockID      string
	latestKDInfo KDResult
}

func saveKDValueToSheet(stockData []KDStockInfo, sheetName pricePeriod) {
	//日期	股票	收盤價	最高價	最低價	RSV	K	D
	//A     B       C      D      E       F   G   H
	sheetData := sheets.ValueRange{
		Range:          fmt.Sprintf("%v!A2:H", sheetName),
		MajorDimension: "ROWS",
	}
	for _, oneStock := range stockData {
		rowData := []interface{}{
			oneStock.latestKDInfo.Date,
			oneStock.stockID,
			oneStock.latestKDInfo.ClosePrice,
			oneStock.latestKDInfo.NHighPrice,
			oneStock.latestKDInfo.NLowPrice,
			oneStock.latestKDInfo.RSV,
			oneStock.latestKDInfo.K,
			oneStock.latestKDInfo.D,
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
