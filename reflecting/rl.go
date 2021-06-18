package reflecting

import (
	"fmt"
	"reflect"
)

// to mimic some specific form of Generics

func printReflect(v interface{}) {
	fmt.Print("Reflect: ")
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Slice {
		return
	}
	for i := 0; i < val.Len(); i++ {
		fmt.Print(val.Index(i).Interface(), " ")
	}
	fmt.Print("\n")
}

func DummyMain() {
	numbers := []int{1, 2, 3}
	strings := []string{"A", "B", "C"}
	floats := []float64{1.5, 2.9, 3.1}
	printReflect(numbers)
	printReflect(strings)
	printReflect(floats)
	printReflect("Not a slices")
}
