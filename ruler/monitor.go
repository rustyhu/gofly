package ruler

// 数据流监控框架
// data monitoring framework

import (
	"log"
)

// Monitor represents the main process of monitoring
type Monitor struct {
	pendingTrans Transaction
	rules        []*Rule
}

func NewMonitor() Monitor {
	const RULE_CAPACITY = 128

	return Monitor{rules: make([]*Rule, 0, RULE_CAPACITY)}
}

func (mt *Monitor) AssignRule(r *Rule) {
	mt.rules = append(mt.rules, r)
}

func (mt *Monitor) InputTransaction(tr Transaction) {
	if tr == nil {
		log.Println("Invalid inserted order!")
		return
	}
	mt.pendingTrans = tr
	mt.update()
	// need external interfere?
	mt.Check()

}

func (mt *Monitor) update() {
	for _, r := range mt.rules {
		r.UpdateData(mt.pendingTrans)
	}
}

func (mt *Monitor) Check() {
	checkChan := make(chan CheckResult, len(mt.rules))
	for _, r := range mt.rules {
		go func(r *Rule) {
			checkChan <- r.Check()
		}(r)
	}

	for i := 0; i < len(mt.rules); i++ {
		if res := <-checkChan; res.Violated {
			// trigger an alert then abandon left results
			mt.alertProc(res.AlertInfo)
			break
		}
	}
}

func (mt *Monitor) alertProc(a *Alarm) {
	if a == nil {
		log.Println("Empty alarm info! Should be a bug!")
		return
	}
	// Raw
	log.Println(*a)
}

func DemoTest() {
	riskctrl := Monitor{}
	for i, entrust := range []Transaction{
		&SimpleOrder{No: 1, Price: 10},
		&SimpleOrder{2, 10, 100, Sell},
	} {
		log.Println("Input entrust no:", i)
		riskctrl.InputTransaction(entrust)
	}
}
