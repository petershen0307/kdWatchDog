package main

import (
	"github.com/petershen0307/kdWatchDog/service"
)

func main() {
	s := service.Scheduler{
		Jobs: []service.ScheduleJob{
			service.GetDailyJob(),
			service.GetWeeklyJob(),
			service.GetMonthlyJob(),
		},
		NumberOfJobWorker: 1,
	}
	s.Run()
}
