package stdlibdemo

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func TestSlicing(t *testing.T) {
	a := []int{0, 2, 4, 6}
	cases := []struct {
		desc string
		data []int
	}{
		{"simple one", []int{-99}},
		{"short", []int{3, 10, 20}},
		{"long", []int{3, 10, 20, 10, 20, 10, 20, 10, 20}},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			passSlice(a, tc.data)
			if len(a) != 4 {
				t.Error("a changed!")
			}
			for i, v := range []int{0, 2, 4, 6} {
				if v != a[i] {
					t.Error("a changed!")
				}
			}
		})
	}
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func isEven(i int) bool {
	return i&1 == 0
}

func TestFilterMut(t *testing.T) {
	isOdd := func(i int) bool {
		return i&1 == 1
	}

	cases := []struct {
		name   string
		before []int
		pred   func(int) bool
		after  []int
	}{
		{"", []int{3, 20, 10}, isOdd, []int{3}},
		{"", []int{3, 20, 10}, isEven, []int{20, 10}},
		{"multiple of 3", []int{-3, 20, 10, 999}, func(i int) bool {
			return i%3 == 0
		}, []int{-3, 999}},
	}

	check := func(a, b []int) bool {
		res := true
		if len(a) != len(b) {
			return false
		}
		for i, v := range a {
			if v != b[i] {
				res = false
				break
			}
		}
		return res
	}

	for _, tt := range cases {
		ttname := tt.name
		if ttname == "" {
			ttname = fmt.Sprint(tt.before, getFunctionName(tt.pred))
		}

		t.Run(ttname, func(t *testing.T) {
			FilterMut(&tt.before, tt.pred)
			if !check(tt.before, tt.after) {
				t.Error("Compare:", tt.before, tt.after)
			}
		})
	}
}
