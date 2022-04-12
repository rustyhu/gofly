package ctrlsys

// 结构化数据流监控
// structured data stream monitor

import (
	"log"
)

// ExamineControl represent the main control process for some type of data, for example risk data
type ExamineControl struct {
	pendingTrans Transaction
	rules        []*Rule
}

func (ctrl *ExamineControl) AssignRule(r *Rule) {
	ctrl.rules = append(ctrl.rules, r)
}

func (ctrl *ExamineControl) InputTransaction(tr Transaction) {
	if tr == nil {
		log.Println("Invalid inserted order!")
		return
	}
	ctrl.pendingTrans = tr
	ctrl.update()
	// need external interfere?
	ctrl.CheckForbidden()

}

func (ctrl *ExamineControl) update() {
	for _, r := range ctrl.rules {
		r.UpdateData(ctrl.pendingTrans)
	}
}

func (ctrl *ExamineControl) CheckForbidden() {
	checkChan := make(chan CheckResult, len(ctrl.rules))
	for _, r := range ctrl.rules {
		go func(r *Rule) {
			checkChan <- r.Check()
		}(r)
	}

	for i := 0; i < len(ctrl.rules); i++ {
		if res := <-checkChan; res.Violated {
			// trigger an alert then abandon left results
			ctrl.alertProc(res.AlertInfo)
			break
		}
	}
}

func (ctrl *ExamineControl) alertProc(a *Alarm) {
	if a == nil {
		log.Println("Empty alarm info! Should be a bug!")
		return
	}
	// Raw
	log.Println(*a)
}

func DemoTest() {
	riskctrl := ExamineControl{}
	for i, entrust := range []Transaction{
		StockOrder{orderBase{No: 1, Price: 10}},
		StockOrder{orderBase{2, 10, 100}},
	} {
		log.Println("Input entrust no:", i)
		riskctrl.InputTransaction(entrust)
	}
}
