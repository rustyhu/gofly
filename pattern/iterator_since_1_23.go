package pattern

import (
	"fmt"
	"iter"
)

// Go 1.23 new feature: range over function (iterator protocol)
// Backward returns an iterator
func Backward[E any](s []E) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

func Backward2[E any](s []E) func(func(string) bool) {
	return func(yield func(string) bool) {
		for i := len(s) - 1; i >= 0; i -= 2 {
			res := fmt.Sprintf("Yield: Idx[%v] - [%v]", i, s[i])
			if !yield(res) {
				return
			}
		}
	}
}

func B3() func(func(string) bool) {
	s := []string{"k", "j", "o", "p", "q"}
	return func(yield func(string) bool) {
		for i := len(s) - 1; i >= 0; i -= 2 {
			res := fmt.Sprintf("Yield: Idx[%v] - [%v]", i, s[i])
			if !yield(res) {
				return
			}
		}
	}
}

func B4(yield func(int, string) bool) {
	s := []string{"k", "j", "o", "p", "q"}
	for i := len(s) - 1; i >= 0; i -= 2 {
		res := fmt.Sprintf("Yield: Idx[%v] - [%v]", i, s[i])
		if !yield(i, res) {
			return
		}
	}
}

// Typical adapter case, based on iterators
func Filter[V any](f func(V) bool, s iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range s {
			if f(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func showCase() {
	// s := []string{"a", "b", "c", "d", "e", "f", "g"}
	// for v := range Backward2(s) {}

	// for v := range B3() {
	// 	fmt.Println(v)
	// }

	for i, v := range B4 {
		fmt.Println(i, v)
	}
}
