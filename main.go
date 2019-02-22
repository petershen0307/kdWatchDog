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
				for _, id := range stockList {
					rawData, err := core.GetStockInfoFromWeb(id, core.WeekPricePeriod)
					if err != nil {
						return
					}
					r := core.KDCalculator(rawData.PriceInfo, 9)
				}
			}},
		},
		NumberOfJobWorker: 1,
	}
	s.Run()
}
