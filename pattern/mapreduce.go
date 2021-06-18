package pattern

import (
	"fmt"
)

type arrIter struct {
	orig []interface{}
}

// map
func (ai *arrIter) forEach(f func(interface{})) {
	for _, ele := range ai.orig {
		f(ele)
	}
}

func (ai *arrIter) filter(f func(interface{}) bool) []interface{} {
	ret := make([]interface{}, 0, len(ai.orig))
	for _, ele := range ai.orig {
		if f(ele) {
			ret = append(ret, ele)
		}
	}
	return ret
}

// reduce
func (ai *arrIter) reduce(f func(interface{}, interface{}) interface{}) interface{} {
	var ret interface{}
	for _, ele := range ai.orig {
		ret = f(ret, ele)
	}
	return ret
}

func RunFunctional() {
	aorig := []interface{}{0, 3, -6, "10", 8}

	arrit := arrIter{aorig}
	for _, i := range arrit.filter(func(ele interface{}) bool {
		k, _ := ele.(int)
		return k&1 != 0
	}) {
		fmt.Println("Get ele: ", i)
	}
}
