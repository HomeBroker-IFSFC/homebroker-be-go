package entity

type OrderType string

const (
	BUY  OrderType = "BUY"
	SELL OrderType = "SELL"
)

func (s OrderType) String() string {
	switch s {
	case BUY:
		return "BUY"
	case SELL:
		return "SELL"
	}
	return "unknown"
}

type OrderStatus string

const (
	OPEN    OrderStatus = "OPEN"
	CLOSED  OrderStatus = "CLOSED"
	PENDING OrderStatus = "PENDING"
	FAILED  OrderStatus = "FAILED"
)

func (s OrderStatus) String() string {
	switch s {
	case OPEN:
		return "OPEN"
	case CLOSED:
		return "CLOSED"
	case PENDING:
		return "PENDING"
	case FAILED:
		return "FAILED"
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
		Status:        OPEN,
		Transactions:  []*Transaction{},
	}
}
