package concurrently

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func tmp(tid int, c rune) int {
	fmt.Printf("This task #%d prints: \"%c\";\n", tid, c)
	// for testcase
	return tid
}

// unit test
func TestRunPool(t *testing.T) {
	expectedIDs := []int{}
	colIDs := struct {
		m   sync.Mutex
		ids map[int]bool
	}{
		ids: make(map[int]bool),
	}

	pool := NewRoutinePool(5, 10)
	id := 0
	for c := 'a'; c <= 'z'; c++ {
		id++
		expectedIDs = append(expectedIDs, id)

		copyID := id
		copyC := c
		pool.PushJob(func() {
			retid := tmp(copyID, copyC)
			colIDs.m.Lock()
			defer colIDs.m.Unlock()
			colIDs.ids[retid] = true
		})
	}
	pool.CloseJobs()
	pool.Wait()

	// number of collected distinct ids equals original ids
	assert.Equal(t, len(expectedIDs), len(colIDs.ids))
}
