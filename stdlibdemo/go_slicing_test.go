package stdlibdemo

import (
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
