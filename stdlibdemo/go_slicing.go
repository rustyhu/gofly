package stdlibdemo

import (
	"fmt"
)

// https://coolshell.cn/articles/21128.html#Slice

func ExamSlicing() {
	arrSmall := make([]int, 10)
	nrArr := arrSmall[:6]

	checkPrint := func(prefix string) {
		fmt.Printf("%v arrSmall addr: %p, %v\n", prefix, arrSmall, arrSmall)
		fmt.Printf("%v nrArr addr: %p, %v\n", prefix, nrArr, nrArr)
	}
	checkPrint("Original")

	// still not beyond capacity of underlying array, extend the slicing range directly
	fmt.Println("Append nrArr, just extend...")
	anum := 3
	nrArr = append(nrArr, anum)
	checkPrint("Now")

	// beyond capacity, allocate a new blocks of memory
	// `arrSmall` start to differ from `nrArr`
	// fmt.Println("Append more, allocating a new block of memory...")
	// arrSmall = append(arrSmall, []int{3, 9, 6, 8}...)
	// checkPrint("Now")

	fmt.Println("Append nrArr beyond limit...")
	nrArr = append(nrArr, []int{99, 99, 99, 99, 199}...)
	checkPrint("Now")

	// do slicing deleting
	fmt.Println("Shrink allSmall...")
	arrSmall = arrSmall[:1]
	checkPrint("Last")
}
