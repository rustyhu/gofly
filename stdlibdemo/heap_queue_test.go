package stdlibdemo

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	// key: price value: List of TradeBlance
	innerQ := &myPriorityQueue{
		&TradeBlance{1, 30, Buy},
		&TradeBlance{2, 12, Buy},
		&TradeBlance{3, 1, Buy},
		&TradeBlance{4, 12, Buy},
	}
	show := func(prefix string) {
		fmt.Println(prefix)
		for _, tb := range *innerQ {
			fmt.Println(*tb)
		}
	}

	show("At beginning:")

	heap.Init(innerQ)
	// fmt.Println("Total:", *innerQ)
	show("After heap init:")

	heap.Push(innerQ, &TradeBlance{8, 30, Sell})
	show("One Push:")

	for i := 0; i < 2; i++ {
		heap.Pop(innerQ)
	}
	show("After 2 Pop:")
}
