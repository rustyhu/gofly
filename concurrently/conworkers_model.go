// model:
// a usually miniature representation of something;
// a pattern of something to be made

package concurrently

import (
	"fmt"
	"sync"
	"time"
)

// Job structure demo
type Job struct {
	id   int
	task func()
}

// Result structure demo
type Result struct {
	info int
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j.id)

		// some heavy work
		j.task()
		time.Sleep(time.Second)

		fmt.Println("worker", id, "finished job", j.id)
		results <- Result{j.id * 2}
	}
}

func ConcurWork() {
	const WORKERS_NUM = 2
	const JOBS_NUM = 80

	var wg sync.WaitGroup
	// 1. Define Job and Result structure
	// the order/sequence of the contents of these channels should be ignored!
	// (do not matter!)
	jobs := make(chan Job, 100)
	results := make(chan Result, 100)

	// 2. Start workers
	// critical region of main routine stepping into concurrent context
	wg.Add(WORKERS_NUM)
	for i := 0; i < WORKERS_NUM; i++ {
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results)
		}(i)
	}

	// 3. Pour all the jobs in (sync or async? that's a choice)
	go func() {
		for i := 0; i < JOBS_NUM; i++ {
			// do nothing in Job.task as a demo
			jobs <- Job{i, func() {}}
		}
		close(jobs)
	}()

	// 4. Final: Wait for finish, processing the results
	// wg.Wait() (sync or async? that's a choice)
	go func() {
		wg.Wait()
		close(results)
	}()

	for rst := range results {
		fmt.Println("Get result:", rst.info)
	}
}
