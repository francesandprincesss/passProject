package job

import "sync"

var wg sync.WaitGroup

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

func (manager *JobManager) Worker(w *WorkerPool) {
	for id := range w.Ch {
		manager.EditJob(id)
		w.wg.Done()
	}
}

func (w *WorkerPool) StartWorker(manager *JobManager) {
	for i := 1; i <= w.WorkerCount; i++ {
		go manager.Worker(w)
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