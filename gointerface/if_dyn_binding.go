package gointerface

// Interface dynamic binding

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64      //计算面积
	perimeter() float64 //计算周长
}

//长方形
type rect struct {
	width, height float64
}

func (r rect) area() float64 { //面积
	return r.width * r.height
}

func (r rect) perimeter() float64 { //周长
	return 2 * (r.width + r.height)
}

//圆形
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func UseIfBinding() {
	r := rect{width: 2.9, height: 4.8}
	c := circle{radius: 4.3}

	shapes := []shape{&r, &c}

	for _, sh := range shapes {
		fmt.Println(sh)
		fmt.Println(sh.area())
		fmt.Println(sh.perimeter())
	}
}
