package main

import (
	"jobQeue/internal/job"
	"time"
)

func main() {

	mainManager := job.NewManager()

	mainPool := job.NewPool(mainManager)

	manager := &job.JobManager{}

	id := mainManager.NewJob("first task")
	mainPool.GetID(id)

	mainPool.StartWorker(manager)

	defer mainPool.WorkerWait()

	time.Sleep(2 * time.Second)
	mainManager.GetAllJob()
	time.Sleep(7 * time.Second)
	mainManager.GetAllJob()

}