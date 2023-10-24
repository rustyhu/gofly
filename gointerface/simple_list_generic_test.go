package gointerface

import (
	"fmt"
	"testing"
)

func TestSimLt(t *testing.T) {
	sl := NewSim[int]()
	sl.Push(3)
	sl.Push(6)
	sl.Push(30)
	v, ok := sl.PopHead()
	fmt.Println("Get a pophead:", v, ok)
	show(sl.Traverse())

	type Record struct {
		id   int
		name string
		// level int
		score int64
	}
	largeL := NewSim[Record]()
	largeL.Push(Record{name: "Zenyen"})
	largeL.Push(Record{id: 1, name: "lawyer", score: 102})
	largeL.Push(Record{id: 2, name: "ancient maga", score: 200})
	nv, ok := largeL.PopHead()
	fmt.Println("Get a pophead:", nv, ok)
	show(largeL.Traverse())
}
