package main

import (
	"fmt"
	"time"

	"github.com/petershen0307/kdWatchDog/service"
)

func main() {
	s := service.Scheduler{
		Jobs: []service.ScheduleJob{
			{JobPeriod: time.Second, JobTriggerTime: time.Now(), JobWork: func() {
				fmt.Println("current time: ", time.Now())
			}},
		},
		NumberOfJobWorker: 1,
	}
	s.Run()
}
