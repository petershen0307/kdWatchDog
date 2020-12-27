package service

import (
	"log"

	"github.com/petershen0307/kdWatchDog/core"
)

func updateKDInfoByPeriod(period core.PricePeriod) {
	// stock list
	stockList := core.GetStockWatchList()
	allStockDailyKD := []core.KDStockInfo{}
	for _, id := range stockList {
		rawData, err := core.GetStockInfoFromWeb(id, period)
		if err != nil {
			log.Printf("Can't get stock(%v) from web. %v", id, err)
			continue
		}
		// skip first 3 data, like yahoo
		r := core.KDCalculator(rawData.PriceInfo[3:], 9)
		allStockDailyKD = append(allStockDailyKD,
			core.KDStockInfo{
				StockID:      id,
				LatestKDInfo: r[len(r)-1],
				StockName:    rawData.Mem.Name,
			})
	}
	core.SaveKDValueToSheet(allStockDailyKD, period)
}
