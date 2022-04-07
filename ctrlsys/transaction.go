package ctrlsys

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
type ETFOrder struct {
	No    uint32
	Price int32
}

///// Transaction END
