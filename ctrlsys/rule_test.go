package ctrlsys

import (
	"testing"
	// "github.com/stretchr/testify/assert"
)

type mockFactorA struct {
	val int64
}

func (f *mockFactorA) Calc() int64 {
	return f.val
}

func (f *mockFactorA) UpdateData(tr Transaction) {
	f.val += tr.getTotalAmount()
}

func TestElemOP(t *testing.T) {
	// assert := assert.New(t)

	mockRules := []struct {
		rule Rule
		res  bool
	}{
		{
			Rule{RuleInfo{MoreEqual, 30}, []Factor{&mockFactorA{29}}, []Operation{}},
			false,
		},
		{
			Rule{RuleInfo{MoreEqual, 30}, []Factor{&mockFactorA{30}}, []Operation{}},
			true,
		},
	}

	for _, r := range mockRules {
		result := r.rule.Check()

		// assert result.Violated is false
		if result.Violated != r.res {
			t.Error("Get rule check result:", result, result.AlertInfo.msg)
		}
	}
}
