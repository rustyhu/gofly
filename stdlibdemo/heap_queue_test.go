package stdlibdemo

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	// key: price value: List of TradeBlance
	rawQ := &myPriorityQueue{
		&Transaction{1, 30, Buy},
		&Transaction{2, 12, Buy},
		&Transaction{3, 1, Buy},
		&Transaction{4, 12, Buy},
	}
	show := func(prefix string) {
		fmt.Println(prefix)
		for _, tb := range *rawQ {
			fmt.Println(*tb)
		}
	}

	show("At beginning:")

	heap.Init(rawQ)
	// fmt.Println("Total:", *innerQ)
	show("After heap init:")

	heap.Push(rawQ, &Transaction{8, 30, Sell})
	show("One Push:")
	// top value is the smallest
	tone := (*rawQ)[0]
	assert.Equal(t, tone.price, 1)

	tone, _ = heap.Pop(rawQ).(*Transaction)
	assert.Equal(t, tone.price, 1)
	tone, _ = heap.Pop(rawQ).(*Transaction)
	assert.Equal(t, tone.price, 12)

	show("After 2 Pop:")
}
