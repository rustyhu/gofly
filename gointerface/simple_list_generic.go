package gointerface

import (
	"fmt"
)

type SData[T any] struct {
	a    T
	next *SData[T]
}

// a generic simple list
type SimLt[T any] struct {
	head *SData[T]
	top  *SData[T]
}

///// behave like a FIFO queue
func NewSim[T any]() SimLt[T] {
	return SimLt[T]{nil, nil}
}

func (l *SimLt[T]) Top() (v T, ok bool) {
	if l.top != nil {
		v, ok = l.top.a, true
	}
	return
}

func (l *SimLt[T]) Push(val T) {
	node := &SData[T]{val, nil}
	// nil pointer exception check
	if l.head == nil {
		l.head = node
	} else {
		// head != nil && top != nil
		l.top.next = node
	}

	l.top = node
}

func (l *SimLt[T]) PopHead() (v T, ok bool) {
	if l.head != nil {
		v, ok = l.head.a, true
		// delete head
		l.head = l.head.next
	}
	return
}

func (l *SimLt[T]) Traverse() []*T {
	res := []*T{}
	for p := l.head; p != nil; p = p.next {
		res = append(res, &p.a)
	}
	return res
}

func show[T any](d []*T) {
	fmt.Println("Traverse:")
	for i, v := range d {
		fmt.Printf("idx[%v] - %v\n", i, *v)
	}
}
