package main

import (
	"fmt"
	gosemaphore "github.com/michaelwp/go-semaphore"
	"sync"
	"time"
)

func main() {

	// set up the maximum goroutine allowed to run in one time.
	sem := gosemaphore.SemaphoreNew(5)
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			sem.Acquire()
			defer sem.Release()
			longRunningProcess(i)
		}(i)
	}

	wg.Wait()
}

func longRunningProcess(taskID int) {
	fmt.Println(time.Now().Format("15:04:05"), "Running task with ID", taskID)
	time.Sleep(2 * time.Second)
}
