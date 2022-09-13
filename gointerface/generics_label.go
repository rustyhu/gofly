package gointerface

import "fmt"

type simA interface {
	getDetail() string
}

type Complex interface {
	simA
	getArea() float32
}

type Circle struct {
	r float32
}

func (c *Circle) getDetail() string {
	return fmt.Sprint("Circle with r:", c.r)
}
func (c *Circle) getArea() float32 {
	return 3.14 * c.r * c.r
}

type Rectangle struct {
	w, h float32
}

func (r Rectangle) getDetail() string {
	return fmt.Sprint("Rec with w:", r.w, "and h:", r.h)
}
func (r Rectangle) getArea() float32 {
	return r.w * r.h
}

// When T is concrete interface (not "any") it can be checked
func GenShow[T Complex](a []T) {
	for i, ele := range a {
		fmt.Println("line", i, ele.getDetail())
		fmt.Println("Get Area:", ele.getArea())
	}
}

func TestG() {
	GenShow([]*Circle{{3}, {4}})
	GenShow([]Rectangle{{3, 5}, {10, 28}})
}
