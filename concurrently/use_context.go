package concurrently

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func longTicker(ctx context.Context) error {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	cnt := 0

	for {
		select {
		case <-ctx.Done():
			// Operation was canceled
			return ctx.Err()

		case <-ticker.C:
			cnt++
			fmt.Printf("Tick%d, Operation is still running...\n", cnt)
		}
	}
}

func hardwork(cancel context.CancelFunc) {
	fmt.Println("Hardwork with context start...")

	// Seed the random number generator
	rseed := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Generate a random integer between 1 and 6
	randomInt := 1 + rseed.Intn(5)
	// Wait for finishing the work then cancel the context
	time.Sleep(time.Duration(randomInt) * time.Second)
	cancel()
	fmt.Printf("Work of %ds finished and context cancel was called.\n", randomInt)
}

func StartCtx() {
	// Create a context with a timeout
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Release resources when main() exits

	go hardwork(cancel)

	// Start the long-running operation
	err := longTicker(ctx)
	if err != nil {
		fmt.Println("Operation failed:", err)
	} else {
		fmt.Println("Operation completed successfully")
	}
}
