package concurrently

// To wait for multiple goroutines to finish

import (
	"fmt"
	"sync"
	"time"
)

func wgChecker(id int, wg *sync.WaitGroup) {
	// Do not forget to notify waitgroup when finished
	defer wg.Done()

	fmt.Printf("Checker %d starting\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Checker %d done\n", id)
}

// WgMain is a dummy main
func WgMain() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go wgChecker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers finished.")
}
