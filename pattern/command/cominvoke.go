package command

import "fmt"

type Order interface {
	execute()
}

type Broker struct {
	orderList []Order
}

func (b *Broker) takeOrder(order Order) {
	b.orderList = append(b.orderList, order)
}
func (b *Broker) pourOrder(orders ...Order) {
	b.orderList = append(b.orderList, orders...)
}
func (b *Broker) placeOrders() {
	for _, o := range b.orderList {
		o.execute()
	}
	// clear all
	b.orderList = nil
}

func Run() {
	// init a stock hold
	stock := Stock{
		name:     "t-nasq",
		quantity: 30,
	}
	fmt.Printf("STOCK Init info: name[%s], quantity[%d]\n", stock.name, stock.quantity)

	b := Broker{orderList: make([]Order, 64)[:0]}
	b.pourOrder(
		NewSellStock(&stock, 20),
		NewBuyStock(&stock, 100),
		NewSellStock(&stock, 20),
		NewSellStock(&stock, 5),
		NewBuyStock(&stock, 300),
	)

	b.placeOrders()
}
