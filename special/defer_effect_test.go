package special

import (
	"fmt"
	"testing"
)

// defer effect
func TestDefer(t *testing.T) {
	num := 3
	defer fmt.Println("Direct defer:", num) // eval value of parameters (like `num`) inmediately

	defer helper(num, 3, t)

	// Using closure - catch the reference
	defer func(target int) {
		fmt.Println("Closure defer:", num)
		if num != target {
			t.Error("closure")
		}
	}(30)

	num = 30
	fmt.Println("Defer get after this line:")
}

func helper(n, expect int, t *testing.T) {
	fmt.Println("In tmpfunc get n:", n)
	if n != expect {
		t.Error()
	}
}
