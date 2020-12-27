package web

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/petershen0307/kdWatchDog/core"
	"github.com/petershen0307/kdWatchDog/service"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world %v", time.Now().Format(time.RFC1123Z))
}

func kdHandler(w http.ResponseWriter, r *http.Request) {
	requestURL := r.URL.String()
	requestID := strings.LastIndex(r.URL.String(), "/")
	// remove the slash /
	requestFunc := requestURL[requestID+1:]
	switch requestFunc {
	case string(core.DailyPricePeriod):
		j := service.GetDailyJob()
		j.JobWork.Run()
	case string(core.WeekPricePeriod):
		j := service.GetWeeklyJob()
		j.JobWork.Run()
	case string(core.MonthPricePeriod):
		j := service.GetMonthlyJob()
		j.JobWork.Run()
	default:
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "ok")
}

// Main is the http listener main function
func Main(port string) {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/kd/", kdHandler)
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
