// Refer to https://stackoverflow.com/questions/38170852/is-this-an-idiomatic-worker-thread-pool-in-go

package concurrently

import (
	"fmt"
	"sync"
)

// JobUnit is the interface(closure) type of this routine-pool
type JobUnit struct {
	jobId  int
	runner func()
}

type RoutinePool struct {
	workersNum int
	// workers sync
	wg sync.WaitGroup

	jobs chan JobUnit
	// also as the assigned job counter
	jobId int
}

func NewRoutinePool(workersNum, jobsCachedVol int) *RoutinePool {
	p := &RoutinePool{
		workersNum: workersNum,
		jobs:       make(chan JobUnit, jobsCachedVol),
		jobId:      0,
	}
	p.RunWorkers()
	// warning: wg can not be copied, so can only be returned by reference
	return p
}

// worker, or named consumer
func (p *RoutinePool) worker(wid int) {
	defer p.wg.Done()

	for job := range p.jobs {
		fmt.Printf("Worker #%d received job[#%d]\n", wid, job.jobId)
		job.runner()
	}
}

func (p *RoutinePool) RunWorkers() {
	// pooling workers:
	for i := 0; i < p.workersNum; i++ {
		p.wg.Add(1)
		go p.worker(i)
	}
}

// Wait for all jobs done
func (p *RoutinePool) Wait() {
	p.wg.Wait()
}

// PushJob into channel once a time
func (p *RoutinePool) PushJob(job func()) {
	p.jobId++
	p.jobs <- JobUnit{p.jobId, job}
}

// CloseJobs manually
func (p *RoutinePool) CloseJobs() {
	close(p.jobs)
}

// task is a function of free form. Design it as what you want
func task(tid int, c rune) {
	fmt.Printf("This task #%d prints: \"%c\";\n", tid, c)
}

func RunPool() {
	pool := NewRoutinePool(4, 8)

	id := 0
	for c := 'a'; c <= 'z'; c++ {
		id++

		// copy to get the independant status for each job, which would be caught by the closure
		copyID := id
		copyC := c
		pool.PushJob(func() {
			task(copyID, copyC)
		})
	}
	pool.CloseJobs()
	pool.Wait()
}
