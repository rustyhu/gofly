package ctrlsys

import "sync"

///// Rule structure

type Alarm struct {
	msg string
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
type Rule struct {
	info    RuleInfo
	factors []Factor
	ops     []Operation // arithmetic operations
}

func (r *Rule) BasicInfo() RuleInfo {
	return r.info
}

func (r *Rule) UpdateData(tr Transaction) {
	for _, f := range r.factors {
		f.UpdateData(tr)
	}
}

func (r *Rule) Check() CheckResult {
	// Calculating Rule formula: consider only +-*/ now, waiting for extending afterward

	// Only able to calculate formula of this form:
	// calcResults[0] (factor[0]) is an anchor, always as the first operant of the formula
	// then each of following operators should be paired with a factor
	// calculation order is always plain left to right
	// Example:
	// 	(f0 + f1 - f3) * f2
	// 	(f0 * f0 + f2) / f1

	calcResults := r.collectFactors()
	// **Reduce** pattern
	calcVal := calcResults[0]
	for _, op := range r.ops {
		// op.operandIdx should never overflow, should be checked when rule was set
		calcVal = op.optor(calcVal, calcResults[op.factorIdx])
	}

	// TODO: notice to compare only integer value, float (like N% ratio
	// comparison) need to do conversion before and after calculation and
	// and take care of float resolution problem
	violated := r.info.compare(calcVal)
	var pAlarm *Alarm = nil
	if violated {
		pAlarm = r.generateAlarm()
	}
	return CheckResult{violated, pAlarm}
}

func (r *Rule) collectFactors() []int64 {
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

	return calcResults
}

func (r *Rule) generateAlarm() *Alarm {
	return &Alarm{"Generating demo"}
}

// arithmeticElemOP consider only four elementary operations of arithmetic (+-*/) now, waiting for extending afterward
type arithmeticElemOP func(lhs, rhs int64) int64

func OPAdd(lhs, rhs int64) int64 {
	return lhs + rhs
}

func OPMns(lhs, rhs int64) int64 {
	return lhs - rhs
}

func OPMul(lhs, rhs int64) int64 {
	return lhs * rhs
}

func OPDiv(lhs, rhs int64) int64 {
	// TODO float resolution problem
	return lhs / rhs
}

type Operation struct {
	optor     arithmeticElemOP
	factorIdx int
}

type Factor interface {
	Calc() int64
	UpdateData(tr Transaction)
}

// Parameters are with Factor
type Parameter struct {
}

// a crude mock: one type of detailed rule statistic data structure
// type TradeBlance struct {
// 	no    int
// 	price int
// 	dir   EntrustDir
// }
