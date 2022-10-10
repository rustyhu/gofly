package special

import (
	"testing"
)

func minmaxCommon(op func(i int, res *int), a []int) int {
	if len(a) == 0 {
		return 0
	}
	res := a[0]
	for _, i := range a {
		op(i, &res)
	}
	return res
}

func min(a ...int) int {
	return minmaxCommon(
		func(i int, res *int) {
			if i < *res {
				*res = i
			}
		},
		a)
}

func max(a ...int) int {
	return minmaxCommon(
		func(i int, res *int) {
			if i > *res {
				*res = i
			}
		},
		a)
}

func TestVariadicFunction(t *testing.T) {
	testcases := []struct {
		params []int
		res    int
		errMsg string
	}{
		{[]int{4, 2, 3, 1}, 1, "Wrong min()!"},
		{[]int{4, 0, 33, 199}, 0, "Wrong min() - 0"},
		{[]int{44, 0, -33, 199},
			-33,
			"Wrong min() - (-33)"},
	}
	for _, tc := range testcases {
		if min(tc.params...) != tc.res {
			t.Error(tc.errMsg)
		}
	}

	testcases = []struct {
		params []int
		res    int
		errMsg string
	}{
		{[]int{4, 2, 3, 1}, 4, "Wrong max()!"},
		{[]int{-4, 0, -33, -199}, 0, "Wrong max() - 0"},
		{[]int{-44, -100, -33, -199},
			-33,
			"Wrong max() - (-33)"},
	}
	for _, tc := range testcases {
		if max(tc.params...) != tc.res {
			t.Error(tc.errMsg)
		}
	}
}
