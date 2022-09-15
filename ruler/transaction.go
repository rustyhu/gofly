package ruler

// Transaction is the process unit of the data stream waiting for check (like entrust, order, ...)
type Transaction interface {
	getTotalBalance() int64
}

///// Transaction details

type EntrustDir int

const (
	Buy EntrustDir = iota
	Sell
	Open
	Close
	Purchase
	Redeem
)

// ETFOrder is an example of detailed `Transaction`
type SimpleOrder struct {
	No     uint64
	Price  uint64
	Volume uint64
	Dir    EntrustDir
	// maybe other things
}

func (s *SimpleOrder) getTotalBalance() int64 {
	return int64(s.Price * s.Volume)
}
