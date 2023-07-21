package entity

import (	
	"time"
	"github.com/google/uuid"
) 



type Transaction struct {
	ID           string
	SellingOrder *Order
	BuingOrder   *Order
	Shares       int
	Price        float64
	Total        float64
	DateTime     time.Time
}


func NewTransaction (sellingOrder *Order, buingOrder *Order, shares int, price float64) *Transaction {
	total := float64(shares) * price 
	return &Transaction{
		ID: uuid.New().String(),
		SellingOrder: sellingOrder,
		BuingOrder: buingOrder,
		Shares: shares,
		Price: price,
		Total: total,
		DateTime: time.Now(),
	}
}