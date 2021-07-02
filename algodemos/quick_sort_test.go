package algodemos

import "testing"

func checkSliceEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range b {
		if v != a[i] {
			return false
		}
	}
	return true
}

func TestQS(t *testing.T) {
	cases := []struct {
		desc   string
		origin []int
		sorted []int
	}{
		{"Empty", []int{}, []int{}},
		{"Already", []int{2, 5, 10}, []int{2, 5, 10}},
		{"Reverse", []int{10, 8, 7, 6, 4, 2, -5},
			[]int{-5, 2, 4, 6, 7, 8, 10}},
		{"General", []int{10, 8, 3, 6, 9}, []int{3, 6, 8, 9, 10}},
		{"many equals 1", []int{3, 2, 5, 5, 2, 2, 3},
			[]int{2, 2, 2, 3, 3, 5, 5}},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			Qsort(c.origin)
			if !checkSliceEq(c.origin, c.sorted) {
				t.Error("Sort faild!")
			}
		})
	}
}
