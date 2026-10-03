package job

import "fmt"

func (manager *JobManager) Worker() {
	for id := range ch {
		fmt.Println(id)
		manager.EditJob(id)
	}
}