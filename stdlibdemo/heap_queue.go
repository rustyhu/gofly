package stdlibdemo

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

type myPriorityQueue []*TradeBlance

func (q myPriorityQueue) Len() int {
	return len(q)
}

func (q myPriorityQueue) Less(i, j int) bool {
	return q[i].price < q[j].price
}

func (q myPriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *myPriorityQueue) Push(x interface{}) {
	*q = append(*q, x.(*TradeBlance))
}

func (q *myPriorityQueue) Pop() interface{} {
	n := len(*q) - 1
	if n < 0 {
		return nil
	}
	popone := (*q)[n]
	(*q)[n] = nil // avoid memory leak

	*q = (*q)[:n]
	return popone
}
