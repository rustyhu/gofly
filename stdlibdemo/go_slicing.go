package stdlibdemo

import (
	"fmt"
)

// https://coolshell.cn/articles/21128.html#Slice

func Sexam() {
	arrSmall := make([]int, 10)
	fmt.Printf("Original array addr: %p\n", arrSmall)
	fmt.Printf("array content: %v\n", arrSmall)
	nrArr := arrSmall[:6]
	fmt.Printf("Original nrArr addr: %p\n", nrArr)

	// still not beyond capacity of underlying array, extend the slicing range directly
	anum := 3
	nrArr = append(nrArr, anum)
	fmt.Printf("Now array content: %v\n", arrSmall)
	fmt.Printf("Now nrArr addr: %p\n", nrArr)

	// beyond capacity, allocate a new blocks of memory
	// `arrSmall` start to differ from `nrArr`
	fmt.Println("Appending more...")
	arrSmall = append(arrSmall, []int{3, 9, 6, 8}...)
	fmt.Printf("Now array addr: %p\n", arrSmall)
	fmt.Printf("array content: %v\n", arrSmall)
	fmt.Printf("Now nrArr addr: %p\n", nrArr)

	// how to do slicing deleting or cleaning?
	arrSmall = arrSmall[:1]
	fmt.Printf("array content: %v\n", arrSmall)
}
