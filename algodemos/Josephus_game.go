package algodemos

import "fmt"

// Josephus Problem

func OneStep(pidx *int, circle []int) {
	circleMove := func(idx *int, l int) {
		(*idx)++
		if *idx == l {
			*idx = 0
		}
	}

	circleMove(pidx, len(circle))
	for circle[*pidx] == -1 {
		circleMove(pidx, len(circle))
	}
}

func JosephusGame(n, k int) int {
	// init
	circle := make([]int, n)
	for i := 0; i < n; i++ {
		circle[i] = i
	}

	idx := 0
	for n > 1 {
		for i := 0; i < k; i++ {
			OneStep(&idx, circle)
		}
		fmt.Println("DEBUG: kill ", circle[idx])
		circle[idx] = -1
		fmt.Println("DEBUG: now ", circle)
		OneStep(&idx, circle)
		n--
	}
	return circle[idx]
}

func Jexam() {
	n, k := 5, 2
	fmt.Printf("N: %d, K: %d\n", n, k)
	fmt.Println("Result:", JosephusGame(n, k))
}
