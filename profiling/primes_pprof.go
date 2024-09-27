package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"time"

	"strconv"
)

// Run this file independently!
//
// Ref: https://go.dev/blog/pprof
// There are 2 methods to run app embeded profiling, with corresponding Go stdlib packages
//   - "runtime/pprof"
//     To collect profiling data into local files.
//   - "net/http/pprof"
//     To start a embeded http server as an interface to provide collected profiling data.
//
// and `go tool pprof` (https://github.com/google/pprof) is another independent
// powerful tool for visualization and analysis of profiling data.

func main() {
	var n int

	if len(os.Args) < 2 {
		n = 800_000
		addr := "localhost:6060"
		go func() {
			log.Println(http.ListenAndServe(addr, nil))
		}()
		log.Printf("Web mode %s, with default primes limit %d...\n", addr, n)

		targetWrapper(n, false)
		// wait for Ctrl+C to quit
		for {
			time.Sleep(time.Second)
		}

	} else {
		// parse 1st argument of commandline as n
		log.Println("Local runtime mode...")
		var err error
		n, err = strconv.Atoi(os.Args[1])
		if err != nil || n <= 0 {
			log.Fatalln("Invalid arg 1 as a prime limit!")
		}

		// Create a file for runtime CPU profiling
		f, err := os.Create("./cpu.pprof")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()

		targetWrapper(n, true)
	}
}

func targetWrapper(n int, bPt bool) {
	// get primes
	primes := primes(n)
	divide3Groups(primes, bPt)
}

func divide3Groups(rawCol []int, bPrint bool) {
	// 100, 1000, 10000
	res := make([][]int, 3)
	for _, i := range rawCol {
		if i < 100 {
			res[0] = append(res[0], i)
		} else if i < 1000 {
			res[1] = append(res[1], i)
		} else {
			res[2] = append(res[2], i)
		}
	}

	if bPrint {
		log.Println("Show -")
		log.Println("Range 1~100:", res[0])
		log.Println("Range 100~1000:", res[0])
		log.Println("Range 1000~:", res[2])
	}
}

// Get all primes less than or equal to n
func primes(n int) []int {
	var primes []int
	for i := 2; i <= n; i++ {
		isPrime := true
		for _, p := range primes {
			if i%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}
