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

type Job struct {
	ID 				int
	TaskName  string
	Status  	JobStatus
	StartTime time.Time
}

type JobManager struct {
	tasks []Job
	mu		sync.RWMutex
}

func NewManager() *JobManager {
	return &JobManager{
		tasks: []Job{},
	}
}

func (manager *JobManager) NewJob(name string) int {
	newID := 1
	manager.mu.Lock()
	for i := range manager.tasks{
		if manager.tasks[i].ID >= newID{
			newID = manager.tasks[i].ID + 1
		}
	}
	job := Job{
		ID: newID,
		TaskName: name,
		Status: StatusPending,
	}
	manager.tasks = append(manager.tasks, job)
	manager.mu.Unlock()
	return newID
}

func (manager *JobManager) FindJob(id int) (*Job, error){
	manager.mu.Lock()
	for n := range manager.tasks {
		if id == manager.tasks[n].ID {
			manager.mu.Unlock()
			return &manager.tasks[n], nil
		}
	}
	manager.mu.Unlock()
	return nil, fmt.Errorf("ID not found")
}

func (manager *JobManager) EditJob(id int) {
	ptr, err := manager.FindJob(id); if err != nil {
		fmt.Println(err)
		return
	}

	manager.mu.Lock()
	ptr.Status = StatusRunning
	ptr.StartTime = time.Now()
	manager.mu.Unlock()

	time.Sleep(5 * time.Second)

	ptr1, err := manager.FindJob(id); if err != nil {
		fmt.Println(err)
		return
	}

	manager.mu.Lock()
	ptr1.Status = StatusCompleted
	manager.mu.Unlock()
}

func (m *JobManager) GetAllJob() {
	m.mu.RLock()
	for i, n := range m.tasks {
		fmt.Println(i+1, "Task")
		fmt.Println(" - - - - - ")
		fmt.Println("ID:", n.ID)
		fmt.Println("Task name:", n.TaskName)
		fmt.Println("Status:", n.Status)
		fmt.Println("Start time:", n.StartTime.Format("02.02.2006 15:04:04"))
		fmt.Println("")
	}
	m.mu.RUnlock()
}