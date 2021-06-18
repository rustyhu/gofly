package algodemos

// Classic backtrace algorithm operation

import "fmt"

func queen(a []int, cur int, pnum *int) {
	if cur == len(a) {
		fmt.Println(a)
		*pnum++
		return
	}
	for r := 0; r < len(a); r++ {
		a[cur] = r
		flag := true
		for j := 0; j < cur; j++ {
			rdistance := r - a[j]
			if rdistance < 0 {
				rdistance = -rdistance
			}
			if rdistance == 0 || rdistance == cur-j {
				flag = false
				break
			}
		}
		if flag {
			queen(a, cur+1, pnum)
		}
	}
}

// Calc `Queen` problem methods and count the total number
func Qexam(l int) {
	rb := make([]int, l)
	num := 0

	queen(rb, 0, &num)
	fmt.Println("number of solves: ", num)
}
