package ctrlsys

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

type RuleA struct {
	info    RuleInfo
	factors []Factor
}

func (r *RuleA) BasicInfo() RuleInfo {
	return r.info
}

func (r *RuleA) Check() CheckResult {
	res := false
	for _, f := range r.factors {
		if f.Check() {
			res = true
			break
		}
	}
	return CheckResult{res, r.generateAlarm()}
}

func (r *RuleA) generateAlarm() *Alarm {
	return &Alarm{}
}

type Factor struct {
}

func (r *Factor) Check() bool {
	return false
}

// a crude mock: one type of detailed rule statistic data structure
// type TradeBlance struct {
// 	no    int
// 	price int
// 	dir   EntrustDir
// }
