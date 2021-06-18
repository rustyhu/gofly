package concurrently

import "fmt"

/*
check memory operation order
https://golang.google.cn/ref/mem
*/

var target string = "initial_value"

func read1(c chan int) {
	fmt.Println("Catch1: ", target)
	c <- 1
}

func read2(c chan int) {
	fmt.Println("Catch2: ", target)
	c <- 1
}

// DummyMain is the executor
func DummyMain() {
	c := make(chan int)
	// actions in the func happened after "go func" call
	go read1(c)
	// channel receiving happened after sending
	<-c

	// this modification happened before read2 running asynchronically
	target = "rewrite 2"
	go read2(c)
	<-c
}
