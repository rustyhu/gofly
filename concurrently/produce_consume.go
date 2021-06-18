package concurrently

// Refer to https://stackoverflow.com/questions/38170852/is-this-an-idiomatic-worker-thread-pool-in-go

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Job struct {
	Id   int
	Work string
}

func produceLine(jobs chan<- *Job) {
	// Generate jobs:
	id := 0
	for c := 'a'; c <= 'z'; c++ {
		id++
		jobs <- &Job{Id: id, Work: fmt.Sprintf("%c", c)}
	}
	close(jobs)
}

func consume(id int, jobs <-chan *Job) {
	defer wg.Done()
	for job := range jobs {
		sleepMs := rand.Intn(1000)
		fmt.Printf("worker #%d received: '%s', sleep %dms\n", id, job.Work, sleepMs)
		time.Sleep(time.Duration(sleepMs) * time.Millisecond)
		fmt.Printf("result: %s\n", job.Work+fmt.Sprintf("-%dms", sleepMs))
	}
}

func Execute() {
	jobs := make(chan *Job, 100) // Buffered channel

	// Start consumers:
	for i := 0; i < 5; i++ { // 5 consumers
		wg.Add(1)
		go consume(i, jobs)
	}
	// Start producing
	go produceLine(jobs)

	wg.Wait() // Wait all consumers to finish processing jobs
}
