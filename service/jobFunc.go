package service

import (
	"time"

	"github.com/petershen0307/kdWatchDog/core"
)

func updateKDInfoByPeriod(period core.PricePeriod) {
	// stock list
	stockList := core.GetStockWatchList()
	allStockDailyKD := []core.KDStockInfo{}
	for _, id := range stockList {
		rawData, err := core.GetStockInfoFromWeb(id, period)
		if err != nil {
			return
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

// GetDailyJob return the daily scheduler job
func GetDailyJob() ScheduleJob {
	return ScheduleJob{
		JobName: "Daily", JobPeriod: 24 * time.Hour, JobTriggerTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 16, 0, 0, 0, time.Local),
		JobWork: func() {
			updateKDInfoByPeriod(core.DailyPricePeriod)
		},
	}
}
