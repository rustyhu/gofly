package ctrlsys

// 大批量结构化数据监控

import "log"

type Dir int

const (
	Buy Dir = iota + 1
	Sell
	Pur
	Red
)

// a crude mock
type TradeBlance struct {
	no    int
	price int
	dir   Dir
}

///// Entrust
type Entrust interface {
}

// example
type ETFEntrustOrder struct {
	No    uint32
	Price int32
}

///// Entrust END

// DataControl represent the main control process for some type of data, for example risk data
type DataControl struct {
	currentOrder Entrust
	rules        []Rule
}

func (ctrl *DataControl) EntrustInput(order Entrust) {
	if order == nil {
		log.Println("Invalid inserted order!")
		return
	}
	ctrl.currentOrder = order
}

func (ctrl *DataControl) AssignRule(r Rule) {
	ctrl.rules = append(ctrl.rules, r)
}

func (ctrl *DataControl) CheckForbidden() {
	for _, r := range ctrl.rules {
		r.Check()
	}
}

func DemoTest() {
	riskctrl := DataControl{}
	for i, entrust := range []Entrust{
		ETFEntrustOrder{1, 10},
		ETFEntrustOrder{2, 10},
		ETFEntrustOrder{3, 20},
	} {
		log.Println("Input entrust no:", i)
		riskctrl.EntrustInput(entrust)
	}
}
