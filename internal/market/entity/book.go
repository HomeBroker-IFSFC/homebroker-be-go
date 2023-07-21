package entity

import (
	"container/heap"
	"sync"
)

type Book struct {
	Order           []*Order
	Transaction     []*Transaction
	OrdersChanel    chan *Order // Input do Kafka
	OrdersChanelOut chan *Order
	Wg              *sync.WaitGroup
}

func NewBook(orderChanel chan *Order, orderChanelOut chan *Order, wg *sync.WaitGroup) *Book {
	return &Book{
		Order:           []*Order{},
		Transaction:     []*Transaction{},
		OrdersChanel:    orderChanel,
		OrdersChanelOut: orderChanelOut,
		Wg:              wg,
	}
}

func (b *Book) Trade() {
	buyOrders := NewOrderQueue()
	sellOrders := NewOrderQueue()

	heap.Init(buyOrders)
	heap.Init(sellOrders)

	for order := range b.OrdersChanel {
		if order.OrderType == Buy {
			buyOrders.Push(order)
			if sellOrders.Len() > 0 && sellOrders.Orders[0].Price <= order.Price {
				sellOrder := sellOrders.Pop().(*Order)
				if sellOrder.PendingShares > 0 {
					// pode negociar
					transaction := NewTransaction(sellOrder, order, order.Shares, sellOrder.Price)
					b.AddTransaction(transaction, b.Wg)
					sellOrder.Transactions = append(sellOrder.Transactions, transaction)
					b.OrdersChanelOut <- sellOrder
					b.OrdersChanelOut <- order

					if sellOrder.PendingShares > 0 {
						sellOrders.Push(sellOrder)
					}
				}
			}
		} else if order.OrderType == Sell {
			sellOrders.Push(order)
			if buyOrders.Len() > 0 && buyOrders.Orders[0].Price >= order.Price {
				buyOrder := buyOrders.Pop().(*Order)
				if buyOrder.PendingShares > 0 {
					transaction := NewTransaction(order, buyOrder, order.Shares, buyOrder.Price)
					b.AddTransaction(transaction, b.Wg)
					buyOrder.Transactions = append(buyOrder.Transactions, transaction)
					order.Transactions = append(order.Transactions, transaction)
					b.OrdersChanelOut <- buyOrder
					b.OrdersChanelOut <- order
					if buyOrder.PendingShares > 0 {
						buyOrders.Push(buyOrder)
					}
				}
			}
		}
	}

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func (b *Book) AddTransaction(transaction *Transaction, wg *sync.WaitGroup) {
	defer wg.Done() // Executa ao fim do método

	sellingShares := transaction.SalesOrder.PendingShares
	buyingShares := transaction.PurchaseOrder.PendingShares

	minShares := min(sellingShares, buyingShares)

	transaction.SalesOrder.Investor.UpdateAssetPosition(transaction.SalesOrder.Asset.ID, -minShares)
	transaction.SalesOrder.PendingShares -= minShares
	transaction.PurchaseOrder.Investor.UpdateAssetPosition(transaction.PurchaseOrder.ID, minShares)
	transaction.PurchaseOrder.PendingShares -= minShares

	transaction.CalculateTotal(transaction.Shares, transaction.Price)

	transaction.CloseBuyingOrder()
	transaction.CloseSellingOrder()

	b.Transaction = append(b.Transaction, transaction)
}
