package command

import "fmt"

// This is the backup mod, usually clean

// defination: a stock info, or a stock hold? more like a stock hold
type Stock struct {
	name     string
	quantity int
}

func (s *Stock) buy(n int) {
	s.quantity += n
	fmt.Println(n, "bought -",
		"Stock [ Name: ", s.name, ",  Quantity: ", s.quantity, " ]")
}
func (s *Stock) sell(n int) {
	s.quantity -= n
	fmt.Println(n, "sold -",
		"Stock [ Name: ", s.name, ", Quantity: ", s.quantity, " ]")
}

type BuyStock struct {
	abcStock *Stock
	n        int
}

func NewBuyStock(s *Stock, n int) BuyStock {
	return BuyStock{
		abcStock: s,
		n:        n,
	}
}
func (bs BuyStock) execute() {
	bs.abcStock.buy(bs.n)
}

type SellStock struct {
	abcStock *Stock
	n        int
}

func NewSellStock(s *Stock, n int) SellStock {
	return SellStock{
		abcStock: s,
		n:        n,
	}
}
func (bs SellStock) execute() {
	bs.abcStock.sell(bs.n)
}
