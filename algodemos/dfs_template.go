package algodemos

import (
	"fmt"
	"sort"
)

// permutation is implemented using DFS template
func Permutation(A []int, recordlen int) [][]int {
	res := make([][]int, 0)
	La := len(A)
	if La == 0 {
		return res
	}

	cur := make([]int, recordlen)
	// used := make([]bool, La)

	// in case of repeat
	sort.Ints(A)
	var arrange func(curidx int)
	arrange = func(curidx int) {
		// recursion stop condition
		if curidx == recordlen {
			rec := make([]int, recordlen)
			copy(rec, cur)
			res = append(res, rec)
			return
		}

		// traversing range (breadth range, as depth first)
		for i := 0; i < La; i++ {
			// in case of repeat
			// if used[i] || (i > 0 && A[i] == A[i-1] && !used[i-1]) {
			// 	continue
			// }
			if i > 0 && A[i] == A[i-1] {
				continue
			}

			// collect
			cur[curidx] = A[i]
			// neccesary record
			// used[i] = true
			// next level
			arrange(curidx + 1)
			// neccesary record recover (backtrace)
			// used[i] = false
		}
	}

	arrange(0)
	return res
}

func PermutationShow() {
	for i, p := range Permutation([]int{2, 3, 30, 99, 2}, 2) {
		fmt.Printf("%dth p: %v\n", i+1, p)
	}
}
