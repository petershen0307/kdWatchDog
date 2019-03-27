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
			{JobPeriod: 24 * time.Hour, JobTriggerTime: time.Now(), JobWork: func() {
				fmt.Println("current time: ", time.Now())
				// stock list
				stockList := []string{"1722", "1726", "2204", "3388", "006208"}
				//stockList := []string{"006208"}
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
						})
				}
				core.SaveKDValueToSheet(allStockDailyKD, core.DailyPricePeriod)
			}},
		},
		NumberOfJobWorker: 1,
	}
	s.Run()
}
