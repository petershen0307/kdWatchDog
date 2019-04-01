package main

import (
	"fmt"
	"time"

	"github.com/petershen0307/kdWatchDog/core"
	"github.com/petershen0307/kdWatchDog/service"
)

func main() {
	s := service.Scheduler{
		Jobs: []service.ScheduleJob{
			{JobPeriod: 24 * time.Hour, JobTriggerTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 16, 0, 0, 0, time.Local), JobWork: func() {
				fmt.Println("current time: ", time.Now())
				// stock list
				stockList := core.GetStockWatchList()
				allStockDailyKD := []core.KDStockInfo{}
				for _, id := range stockList {
					rawData, err := core.GetStockInfoFromWeb(id, core.DailyPricePeriod)
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
				core.SaveKDValueToSheet(allStockDailyKD, core.DailyPricePeriod)
			}},
		},
		NumberOfJobWorker: 1,
	}
	s.Run()
}
