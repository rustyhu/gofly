package gointerface

import "fmt"

type simpleProc struct {
	decorator int
}

func (sp *simpleProc) functor(a, b int) int {
	fmt.Println("functor start...")
	return a + b
}

func (sp *simpleProc) functor2nd(a, b int) int {
	fmt.Println("functor2nd start...")
	fmt.Println("Show decorator:", sp.decorator)
	return a * b
}

func (sp *simpleProc) Caller(a int) int {
	fmt.Println("Caller start...")
	// bind specific func signature to `struct method`
	var innerf func(int, int) int
	if a < 2 {
		innerf = sp.functor
	} else {
		innerf = sp.functor2nd
	}

	fmt.Println("Caller end...")
	return innerf(a, 2)
}

func memain() {
	s := simpleProc{}
	fmt.Println("Final ret:", s.Caller(3))
	var funcVar func(int, int) int = s.functor2nd

	funcVar(30, 60)
}
