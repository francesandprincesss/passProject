package job

import (
	"fmt"
	"sync"
	"time"
)

type JobStatus string

const (
	StatusPending  JobStatus = "pending"
	StatusRunning  JobStatus = "running"
	StatusCompleted  JobStatus = "completed"
)

var ch = make(chan int, 4)

type Job struct {
	ID 				int
	TaskName  string
	Status  	JobStatus
	StartTime time.Time
}

type JobManager struct {
	Tasks []Job
	mu		sync.RWMutex
}

func NewManager() *JobManager {
	return &JobManager{
		Tasks: []Job{},
	}
}

func (manager *JobManager) NewJob(name string) {
	newID := 1
	manager.mu.Lock()
	for i := range manager.Tasks{
		if manager.Tasks[i].ID >= newID{
			newID = manager.Tasks[i].ID + 1
		}
	}
	job := Job{
		ID: newID,
		TaskName: name,
		Status: StatusPending,
	}
	manager.Tasks = append(manager.Tasks, job)
	manager.mu.Unlock()
	ch <- newID
}

func (manager *JobManager) FindJob(id int) (*Job, error){
	manager.mu.Lock()
	for n := range manager.Tasks {
		if id == manager.Tasks[n].ID {
			return &manager.Tasks[n], nil
		}
	}
	manager.mu.Unlock()
	return nil, fmt.Errorf("ID not found")
}

func (manager *JobManager) EditJob(id int) {
	ptr, err := manager.FindJob(id); if err != nil {
		fmt.Println(err)
		fmt.Println("editing1212")
		return
	}

	manager.mu.Lock()
	ptr.Status = StatusRunning
	ptr.StartTime = time.Now()
	manager.mu.Unlock()

	fmt.Println("editing too")
	time.Sleep(5 * time.Second)

	manager.mu.Lock()
	ptr.Status = StatusCompleted
	manager.mu.Unlock()
}