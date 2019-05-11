package service

import (
	"log"
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
			log.Fatalln("Can't get stock from sheet.", err)
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

// GetDailyJob return the daily kd scheduler job
func GetDailyJob() ScheduleJob {
	return ScheduleJob{
		JobName: "Daily", JobPeriod: 24 * time.Hour, JobTriggerTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 16, 0, 0, 0, time.Local),
		JobWork: func() {
			updateKDInfoByPeriod(core.DailyPricePeriod)
		},
	}
}

// GetWeeklyJob return the weekly kd scheduler job
func GetWeeklyJob() ScheduleJob {
	return ScheduleJob{
		JobName: "Weekly", JobPeriod: 2 * 24 * time.Hour, JobTriggerTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 16, 10, 0, 0, time.Local),
		JobWork: func() {
			updateKDInfoByPeriod(core.WeekPricePeriod)
		},
	}
}

// GetMonthlyJob return the monthly kd scheduler job
func GetMonthlyJob() ScheduleJob {
	return ScheduleJob{
		JobName: "Monthly", JobPeriod: 2 * 24 * time.Hour, JobTriggerTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 16, 20, 0, 0, time.Local),
		JobWork: func() {
			updateKDInfoByPeriod(core.MonthPricePeriod)
		},
	}
}
