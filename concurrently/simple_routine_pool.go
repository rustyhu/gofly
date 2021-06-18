// Refer to https://stackoverflow.com/questions/38170852/is-this-an-idiomatic-worker-thread-pool-in-go

package concurrently

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Functor is the interface(closure) type of this routine-pool
type Functor func()

func produce(jobs chan<- Functor) {
	id := 0
	for c := 'a'; c <= 'z'; c++ {
		id++

		// copy to get the independant status, which would be caught by the closure
		copyID := id
		copyC := c
		jobs <- func() {
			fmt.Printf("This job #%d prints: \"%c\";\n", copyID, copyC)
		}
	}

	close(jobs)
}

// worker, or named consumer
func worker(id int, jobs <-chan Functor, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		sleepMs := rand.Intn(100)
		fmt.Printf("Worker #%d received job, suspend for %dms\n", id, sleepMs)
		time.Sleep(time.Duration(sleepMs) * time.Millisecond)
		job()
	}
}

func demoExe() {
	jobs := make(chan Functor, 100)
	// routine pool sync waiter
	var wg sync.WaitGroup

	// pooling workers:
	const NUM = 5
	for i := 0; i < NUM; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Start producing jobs
	go produce(jobs)

	wg.Wait()
}
