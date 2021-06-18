package algodemos

import "fmt"

func permutations(A []int, plen int) [][]int {
	res := [][]int{}
	rec := make([]int, plen)
	permuRes(&res, rec, 0, A, plen)

	return res
}

func permuRes(res *[][]int, rec []int, recidx int, A []int, plen int) {
	if recidx == plen {
		recc := make([]int, len(rec))
		copy(recc, rec)
		*res = append(*res, recc)
		return
	}

	for _, v := range A {
		repeat := false
		for i := 0; i < recidx; i++ {
			if rec[i] == v {
				repeat = true
				break
			}
		}
		if !repeat {
			rec[recidx] = v
			permuRes(res, rec, recidx+1, A, plen)
		}
	}
}

func combination(A []int, plen int) [][]int {
	res := [][]int{}
	for i := 0; i < len(A)-1; i++ {
		for _, v := range A[i+1:] {
			res = append(res, []int{A[i], v})
		}
	}

	return res
}

func Pexam() {
	// algodemos.Jexam()
	// fmt.Println(permutations([]int{1, 2, 3, 4, 5, 6}, 2))
	fmt.Println(combination([]int{1, 2, 3, 4, 5, 6}, 2))
}
