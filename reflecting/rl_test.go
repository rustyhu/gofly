package reflecting

import "testing"

func TestKkk(t *testing.T) {
	testCases := []struct {
		desc    string
		newName string
	}{
		{
			desc:    "tmp test",
			newName: "tcase",
		},
		// ...
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
		})
	}
}
