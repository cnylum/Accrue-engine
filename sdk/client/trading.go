package client

import (
	"context"

	"github.com/cnylum/accrue-engine/sdk/types"
)

// PlaceOrder places a new order.
func (c *Client) PlaceOrder(ctx context.Context, req types.PlaceOrderRequest) (types.Order, error) {
	var order types.Order
	err := c.post(ctx, "/v1/orders", req, &order)
	return order, err
}

// CancelOrder cancels an existing order.
func (c *Client) CancelOrder(ctx context.Context, orderID string) error {
	return c.delete(ctx, "/v1/orders/"+orderID)
}

// GetOrder retrieves an order by ID.
func (c *Client) GetOrder(ctx context.Context, orderID string) (types.Order, error) {
	var order types.Order
	err := c.get(ctx, "/v1/orders/"+orderID, &order)
	return order, err
}
