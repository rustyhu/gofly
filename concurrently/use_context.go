package concurrently

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func StartCtx() {
	Lseconds := rand.Intn(10) + 1
	GroupN := 4
	worktimeN := make([]int, GroupN)
	for i := range GroupN {
		worktimeN[i] = rand.Intn(10) + 1
	}
	fmt.Printf("Limited seconds is %d;\n", Lseconds)
	for i := range GroupN {
		fmt.Printf("Worker %d worktime is %d seconds\n", i, worktimeN[i])
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(Lseconds)*time.Second)
	defer cancel()

	// wg is used for observing/debug
	wg := sync.WaitGroup{}
	done := make(chan string, GroupN)
	for id, worktime := range worktimeN {
		wg.Add(1)
		go ctxworker(ctx, done, id, worktime, &wg)
	}

	select {
	case res := <-done:
		cancel() // finished, notify still working ones
		fmt.Println("Finally get:", res)
	case <-ctx.Done():
		fmt.Println("All workers failed!")
	}

	wg.Wait()
	// concurrently.StartCtx()
}

func ctxworker(ctx context.Context, done chan<- string, id, worktime int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= worktime; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker%d stoped: %s\n", id, ctx.Err())
			return
		default:
			time.Sleep(time.Second)
		}
	}
	done <- fmt.Sprintf("Task completed by worker%d, get mark %d", id, worktime)
}
