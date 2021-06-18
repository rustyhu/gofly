package concurrently

// pipeline pattern
// https://coolshell.cn/articles/21228.html

// Original echo sig: type Echo func([]int) <-chan int
type EchoFunc func() <-chan int
type PipeFunc func(<-chan int) <-chan int

func Echo(rgmin, rgmax int) <-chan int {
	out := make(chan int)
	go func() {
		for i := rgmin; i <= rgmax; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func pipeline(echo EchoFunc, pipeFns ...PipeFunc) <-chan int {
	ch := echo()
	for i := range pipeFns {
		ch = pipeFns[i](ch)
	}
	return ch
}
