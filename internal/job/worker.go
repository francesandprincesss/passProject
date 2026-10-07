package job

import "sync"

var wg sync.WaitGroup

type Manager interface {
	EditJob(id int)
}

type WorkerPool struct {
	Manager				*JobManager
	WorkerCount   int
	Ch 						chan int
	wg 						sync.WaitGroup

}

func NewPool(manager *JobManager) *WorkerPool {
	return &WorkerPool{
		Manager: 			manager,
		WorkerCount:  3,
		Ch: 					make(chan int, 67),
	}
}

func (w *WorkerPool) Worker(manager Manager) {
	for id := range w.Ch {
		manager.EditJob(id)
		w.wg.Done()
	}
}

func (w *WorkerPool) StartWorker(manager Manager) {
	for i := 1; i <= w.WorkerCount; i++ {
		go w.Worker(manager)
	}
}

func (w *WorkerPool) GetID(id int) {
	w.Ch <- id
	w.wg.Add(1)
}

func (w *WorkerPool) WorkerWait() {
	w.wg.Wait()
	//close(w.Ch)
}