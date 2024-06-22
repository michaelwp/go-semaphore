# go-semaphore
go-semaphore is a Go package that implements the semaphore pattern than can be used to manage the number of Goroutine 
to run concurrently.

### installation
```shell
go get -d github.com/michaelwp/go-semaphore
```

### basic of use

```go
package main

import (
	"fmt"
	gosemaphore "github.com/michaelwp/go-semaphore"
	"time"
)

func main() {
	sem := gosemaphore.SemaphoreNew(5) // <-- set up the maximum allowed number of goroutine to run concurrently.
	
	for i := 0; i < 50; i++ {
		go func(i int) {
			sem.Acquire()       // <-- to acquire a seat in buffered channel. This has to be at the beginning of the process.
			defer sem.Release() // <-- to release the seat. Strongly suggested to use defer to have a guarantee that this function will be executed when the process is finishes.
			Process(i)
		}(i)
	}
}

func Process(taskID int) {
	fmt.Println(time.Now().Format("15:04:05"), "Running task with ID", taskID)
	time.Sleep(2 * time.Second)
}
```
