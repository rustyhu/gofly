package ctrlsys

import "fmt"

///// Transaction is the process unit of the data stream waiting for checking (like entrust, order, ...)
type Transaction interface {
	getSum() int64
}

type EntrustDir int

const (
	Buy EntrustDir = iota + 1
	Sell
	Pur
	Red
)

// ETFOrder is an example of detailed `Transaction`
type SimpleOrder struct {
	No     uint64
	Price  uint64
	Volume uint64
	// securityCode string
	// other things
}

func (s *SimpleOrder) getSum() int64 {
	return int64(s.Price * s.Volume)
}

func demoT() {
	a := SimpleOrder{}
	fmt.Println(a.No, a.Price)
}
