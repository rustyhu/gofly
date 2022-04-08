package ctrlsys

import "sync"

///// Rule structure
type Alarm struct {
}

type CheckResult struct {
	Violated  bool
	AlertInfo *Alarm
}

type RuleDirection int

const (
	Equal RuleDirection = iota
	More
	MoreEqual
	Less
	LessEqual
)

type RuleInfo struct {
	Dir       RuleDirection
	Threshold int64
}

func (ri *RuleInfo) compare(calcRes int64) bool {
	switch ri.Dir {
	case Equal:
		return calcRes == ri.Threshold
	case More:
		return calcRes > ri.Threshold
	case MoreEqual:
		return calcRes >= ri.Threshold
	case Less:
		return calcRes < ri.Threshold
	case LessEqual:
		return calcRes <= ri.Threshold
	default:
		return false
	}
}

// Rule is an info holder, a data holder, a checker, ... may be the key role in this system
type Rule interface {
	// holder
	BasicInfo() RuleInfo
	UpdateData(Transaction)

	// checker
	Check() CheckResult
	generateAlarm() *Alarm
}

///// Rule structure END

type DemoRuleMulitFactors struct {
	info    RuleInfo
	factors []Factor
	ops     []Operation // arithmetic operations
}

func (r *DemoRuleMulitFactors) BasicInfo() RuleInfo {
	return r.info
}

func (r *DemoRuleMulitFactors) Check() CheckResult {
	// Rule formula, consider only +-*/ now, waiting for extending afterward
	calcResults := make([]int64, len(r.factors))
	wg := sync.WaitGroup{}
	for i := range r.factors {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			calcResults[i] = r.factors[i].Calc()
		}(i)
	}
	wg.Wait()

	// calcResults[0] (factor[0]) is an anchor, always first operant of the formula
	res := calcResults[0]
	for _, op := range r.ops {
		// op.operandIdx should never overflow, should be checked when rule was set
		res = op.optor(res, calcResults[op.factorIdx])
	}

	return CheckResult{r.info.compare(res),
		r.generateAlarm()}
}

func (r *DemoRuleMulitFactors) generateAlarm() *Alarm {
	return &Alarm{}
}

// arithmeticElemOP consider only four elementary operations of arithmetic (+-*/) now, waiting for extending afterward
type arithmeticElemOP func(lhs, rhs int64) int64

type Operation struct {
	optor     arithmeticElemOP
	factorIdx int
}

// Parameters are with Factor
type Parameter struct {
}

type Factor struct {
	parameters map[int]Parameter
}

func (r *Factor) Calc() int64 {
	return 0
}

// a crude mock: one type of detailed rule statistic data structure
// type TradeBlance struct {
// 	no    int
// 	price int
// 	dir   EntrustDir
// }
