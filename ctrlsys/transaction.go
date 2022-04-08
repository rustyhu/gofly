package ctrlsys

import "fmt"

///// Transaction is the process unit of the data stream waiting for checking (like entrust, order, ...)
type Transaction interface {
}

type EntrustDir int

const (
	Buy EntrustDir = iota + 1
	Sell
	Pur
	Red
)

// ETFOrder is an example of detailed `Transaction`
type orderBase struct {
	No     uint32
	Price  uint32
	Volume uint32
	// securityCode string
}

type StockOrder struct {
	orderBase
}

func demoT() {
	a := StockOrder{}
	fmt.Println(a.No, a.Price)
}
