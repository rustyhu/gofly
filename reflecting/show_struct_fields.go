package reflecting

import (
	"fmt"
	"reflect"
)

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   int
}

type T struct {
	A int
	B string
}

func showStruct() {
	SomeEe := Employee{}
	// t := T{23, "skidoo"}

	s := reflect.ValueOf(&SomeEe).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, field.Type(), field.Interface())
	}
}
