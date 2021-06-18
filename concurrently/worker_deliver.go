package concurrently

// Refer to https://stackoverflow.com/questions/38170852/is-this-an-idiomatic-worker-thread-pool-in-go

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func consumer(id string, work string, o chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	sleepMs := rand.Intn(1000)
	fmt.Printf("consumer '%s' received: '%s', sleep %dms\n", id, work, sleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	o <- work + fmt.Sprintf("-%dms", sleepMs)
}

func Delivering() {
	var workChannel = make(chan string)
	var resultsChannel = make(chan string)

	// create goroutine per item in work channel,
	// then write results into results Channel.
	go func() {
		var c = 1
		var wg sync.WaitGroup
		for work := range workChannel {
			wg.Add(1)
			go consumer(strconv.Itoa(c), work, resultsChannel, &wg)
			c++
		}
		wg.Wait()
		fmt.Println("Closing results channel...")
		close(resultsChannel)
	}()

	// add work to the work channel
	go func() {
		fmt.Println("Sent work to work channel:")
		for c := 'a'; c <= 'z'; c++ {
			workChannel <- fmt.Sprintf("%c", c)
		}
		close(workChannel)
	}()

	for x := range resultsChannel {
		fmt.Printf("result: %s\n", x)
	}
}
