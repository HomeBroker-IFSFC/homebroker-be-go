package entity

type OrderType uint8

const (
	Buy OrderType = iota
	Sell
)

func (s OrderType) String() string {
	switch s {
	case Buy:
		return "Buy"
	case Sell:
		return "Sell"
	}
	return "unknown"
}

type OrderStatus uint8

const (
	open OrderStatus = iota
	closed
)

func (s OrderStatus) String() string {
	switch s {
	case open:
		return "open"
	case closed:
		return "close"
	}
	return "unknown"
}

type Order struct {
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	PendingShares int
	Price         float64
	OrderType     OrderType
	Status        OrderStatus
	Transactions  []*Transaction
}

func NewOrder(orderId string, investor *Investor, asset *Asset, shares int, price float64, orderType OrderType) *Order {
	return &Order{
		ID:            orderId,
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		PendingShares: shares,
		Price:         price,
		OrderType:     orderType,
		Status:        open,
		Transactions:  []*Transaction{},
	}
}
