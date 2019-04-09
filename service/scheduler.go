package service

import (
	"log"
	"time"
)

// JobStruct define the job structure
type JobStruct func()

// ScheduleJob is a job configuration
type ScheduleJob struct {
	JobName        string
	JobPeriod      time.Duration
	JobTriggerTime time.Time //Kitchen     = "3:04PM"
	JobWork        JobStruct
}

// Scheduler store many ScheduleJob
type Scheduler struct {
	Jobs              []ScheduleJob
	NumberOfJobWorker uint
	jobQueue          chan JobStruct
}

// unit is ms
const scheduleTimePrecise = time.Second

func (s *Scheduler) workerFunc() {
	for {
		select {
		case x := <-s.jobQueue:
			x()
		default:
			time.Sleep(time.Second)
		}
	}
}

func (s *Scheduler) triggerOnTimeJobs() {
	cTime := time.Now()
	for i, job := range s.Jobs {
		if cTime.After(job.JobTriggerTime) {
			log.Printf("trigger job: %v, time: %v", job.JobName, job.JobTriggerTime)
			s.jobQueue <- job.JobWork
			s.Jobs[i].JobTriggerTime = job.JobTriggerTime.Add(job.JobPeriod)
			log.Printf("job: %v, next trigger time: %v", job.JobName, s.Jobs[i].JobTriggerTime)
		}
	}
}

// Run is the function to execute the Scheduler
func (s *Scheduler) Run() {
	// create the channel as the job queue
	s.jobQueue = make(chan JobStruct, s.NumberOfJobWorker*100)
	// init the worker
	for i := uint(0); i < s.NumberOfJobWorker; i++ {
		go s.workerFunc()
	}
	tick := time.Tick(scheduleTimePrecise)
	for {
		select {
		case <-tick:
			// check scheduler jobs
			s.triggerOnTimeJobs()
		default:
			// go to sleep
			time.Sleep(500 * time.Millisecond)
		}
	}
}
