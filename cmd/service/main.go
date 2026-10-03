package main

import (
	"fmt"
	"jobQeue/internal/job"
	"time"
)

func main() {
	mainManager := job.NewManager()

	mainManager.NewJob("first task")

	fmt.Println(mainManager.Tasks)

	go mainManager.Worker()

	
	time.Sleep(6 * time.Second)
	fmt.Println(mainManager.Tasks)
}