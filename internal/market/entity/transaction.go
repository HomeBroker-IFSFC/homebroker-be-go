package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID            string
	SalesOrder    *Order
	PurchaseOrder *Order
	Shares        int
	Price         float64
	Total         float64
	DateTime      time.Time
}

func NewTransaction(sellingOrder *Order, buingOrder *Order, shares int, price float64) *Transaction {
	total := float64(shares) * price
	return &Transaction{
		ID:            uuid.New().String(),
		SalesOrder:    sellingOrder,
		PurchaseOrder: buingOrder,
		Shares:        shares,
		Price:         price,
		Total:         total,
		DateTime:      time.Now(),
	}
}

func (transaction *Transaction) CloseBuyingOrder() {
	if transaction.PurchaseOrder.PendingShares == 0 {
		transaction.PurchaseOrder.Status = CLOSED
	}
}

func (transaction *Transaction) CloseSellingOrder() {
	if transaction.SalesOrder.PendingShares == 0 {
		transaction.SalesOrder.Status = CLOSED
	}
}

func (transaction *Transaction) CalculateTotal(shares int, price float64) {
	transaction.Total = float64(shares) * price
}
