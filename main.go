package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/petershen0307/kdWatchDog/service"
	"github.com/petershen0307/schedulerGo"
)

func main() {
	s := schedulergo.NewScheduler(1)
	s.AddJob(service.GetDailyJob()).
		AddJob(service.GetWeeklyJob()).
		AddJob(service.GetMonthlyJob())
	s.Run()
	osEvent := make(chan os.Signal, 1)
	signal.Notify(osEvent, os.Interrupt)
	for {
		select {
		case <-osEvent:
			break
		default:
			time.Sleep(time.Second)
		}
	}
	log.Println("Program stop")
	s.Stop()
	// [ToDo] wait scheduler worker stopped
}
