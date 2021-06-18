package main

import (
	"fmt"
	"math"

	"gofly/concurrently"
)

// pipeline pattern
// https://coolshell.cn/articles/21228.html

func PipeFuncTmpl(pipein <-chan int, fn func(int, chan<- int)) <-chan int {
	pipeout := make(chan int)
	go func() {
		for n := range pipein {
			fn(n, pipeout)
		}
		close(pipeout)
	}()
	return pipeout
}

func oddExample(pipein <-chan int) <-chan int {
	return PipeFuncTmpl(pipein, func(i int, c chan<- int) {
		if i%2 == 1 {
			c <- i
		}
	})
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for n := range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

func is_prime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func prime(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if is_prime(n) {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func merge(chans []<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for ich, ch := range chans {
			for n := range ch {
				out <- n
				fmt.Printf("Channel[%v] get a sum: [%v]\n", ich, n)
			}
		}
		close(out)
	}()
	return out
}

func main() {
	in := concurrently.Echo(1, 1000)
	const nP = 5
	var chans [nP]<-chan int

	for i := range chans {
		chans[i] = sum(prime(in))
	}
	for n := range sum(merge(chans[:])) {
		fmt.Println("Total Sum:", n)
	}
}
