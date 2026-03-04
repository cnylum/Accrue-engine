package mock

import (
	"context"
	"fmt"
	"sync"

	"github.com/cnylum/accrue-engine/pkg/id"
	"github.com/cnylum/accrue-engine/services/adapter/internal/broker"
)

// Broker is an in-memory mock broker for paper trading.
type Broker struct {
	mu     sync.RWMutex
	orders map[string]broker.Order
}

// New creates a new mock broker.
func New() *Broker {
	return &Broker{
		orders: make(map[string]broker.Order),
	}
}

func (b *Broker) PlaceOrder(_ context.Context, order broker.Order) (broker.Order, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	order.ID = id.New()
	order.Status = "filled"
	order.FilledQty = order.Quantity
	b.orders[order.ID] = order
	return order, nil
}

func (b *Broker) CancelOrder(_ context.Context, _, orderID string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	o, ok := b.orders[orderID]
	if !ok {
		return fmt.Errorf("order %s not found", orderID)
	}
	o.Status = "cancelled"
	b.orders[orderID] = o
	return nil
}

func (b *Broker) GetOrder(_ context.Context, _, orderID string) (broker.Order, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	o, ok := b.orders[orderID]
	if !ok {
		return broker.Order{}, fmt.Errorf("order %s not found", orderID)
	}
	return o, nil
}

func (b *Broker) GetQuote(_ context.Context, symbol string) (broker.Quote, error) {
	return broker.Quote{
		Symbol: symbol,
		Bid:    "100.00",
		Ask:    "100.50",
		Last:   "100.25",
	}, nil
}
