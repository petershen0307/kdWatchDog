package main

import (
	"github.com/petershen0307/kdWatchDog/service"
)

func main() {
	s := service.Scheduler{
		Jobs: []service.ScheduleJob{
			service.GetDailyJob(),
		},
		NumberOfJobWorker: 1,
	}
	s.Run()
}
