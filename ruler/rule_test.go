package ruler

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
	f.val += tr.getTotalBalance()
}

func TestSingleFactorTotalBalance(t *testing.T) {
	// assert := assert.New(t)

	mockRules := []struct {
		rule *Rule
		res  bool
	}{
		{
			&Rule{RuleInfo{MoreEqual, 30}, []Factor{&mockFactorA{29}}, []Operation{}},
			false,
		},
		{
			&Rule{RuleInfo{MoreEqual, 30}, []Factor{&mockFactorA{30}}, []Operation{}},
			true,
		},
	}

	// Skip the rule update process

	for _, r := range mockRules {
		result := r.rule.Check()

		// assert result.Violated is false
		if result.Violated != r.res {
			t.Error("Get rule check result:", result, result.AlertInfo.msg)
		}
	}

	longRule := Rule{RuleInfo{More, 30}, []Factor{&mockFactorA{0}}, []Operation{}}
	longRule.UpdateData(&SimpleOrder{1, 30, 1, Buy})
	if res := longRule.Check(); res.Violated {
		t.Error("Wrong check when not enough!")
	}

	longRule.UpdateData(&SimpleOrder{2, 1, 1, Buy})
	if res := longRule.Check(); !res.Violated {
		t.Error("Wrong check when more!")
	}
}
