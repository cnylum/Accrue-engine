package broker

import "context"

// Order represents a trade order.
type Order struct {
	ID            string
	AccountID     string
	Symbol        string
	Side          string // "buy" or "sell"
	OrderType     string // "market" or "limit"
	Quantity      string
	LimitPrice    string
	FilledQty     string
	Status        string
}

// Quote represents a market quote.
type Quote struct {
	Symbol string
	Bid    string
	Ask    string
	Last   string
}

// Broker defines the interface that all broker adapters must implement.
type Broker interface {
	PlaceOrder(ctx context.Context, order Order) (Order, error)
	CancelOrder(ctx context.Context, accountID, orderID string) error
	GetOrder(ctx context.Context, accountID, orderID string) (Order, error)
	GetQuote(ctx context.Context, symbol string) (Quote, error)
}
