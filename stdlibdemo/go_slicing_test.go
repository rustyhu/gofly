package stdlibdemo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLenAndCap(t *testing.T) {
	a := []int{0, 2, 4, 6}
	assert.Equal(t, 4, len(a))
	assert.Equal(t, 4, cap(a))

	type Ele struct {
		id      int
		checked bool
	}
	b := make([]Ele, 0, 10)
	assert.Equal(t, 0, len(b))
	for i := 99; i >= 90; i-- {
		b = append(b, Ele{id: i})
	}
	assert.Equal(t, 10, len(b))
	assert.Equal(t, 10, cap(b))
}

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
			passAndAppendSlice(a, tc.data)
			assert.Equal(t, 4, len(a))

			for i, v := range []int{0, 2, 4, 6} {
				assert.Equal(t, v, a[i])
			}
		})
	}
}

func TestFilterMut(t *testing.T) {
	isOdd := func(i int) bool {
		return i&1 == 1
	}

	isEven := func(i int) bool {
		return i&1 == 0
	}

	cases := []struct {
		name   string
		before []int
		pred   func(int) bool
		expect []int
	}{
		{"Odd", []int{3, 20, 10}, isOdd, []int{3}},
		{"Even", []int{3, 20, 10}, isEven, []int{20, 10}},
		{
			"Multiple of 3",
			[]int{-3, 20, 10, 999},
			func(i int) bool { return i%3 == 0 },
			[]int{-3, 999},
		},
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
		t.Run(tt.name, func(t *testing.T) {
			FilterMut(&tt.before, tt.pred)
			assert.Truef(t, check(tt.before, tt.expect),
				"expect: %s, actual: %s", tt.expect, tt.before)
		})
	}
}
