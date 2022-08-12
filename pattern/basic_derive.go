package pattern

import "fmt"

// How to share common behavior like OOP base class?

type Base struct {
	// some basic fields ...
	// This is the requirements for "derives"
	proc specProc
}

func (b *Base) Run() {
	fmt.Println("Common part1...")

	if b.proc.init() {
		fmt.Println("Common part3 depending on condition...")
		b.proc.step2()
	}
	b.proc.show()

	fmt.Println("Common part2...")
}

type specProc interface {
	init() bool
	step2()
	show()
}

type Product1 struct {
	val int
}

func (p *Product1) init() bool {
	return p.val == 2
}

func (p *Product1) step2() {
	fmt.Println("Product1 step2!")
}
func (p *Product1) show() {
	fmt.Println("Product1 show:", p.val)
}
