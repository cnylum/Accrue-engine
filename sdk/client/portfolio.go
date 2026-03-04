package client

import (
	"context"

	"github.com/cnylum/accrue-engine/sdk/types"
)

// GetPositions returns all positions for the authenticated account.
func (c *Client) GetPositions(ctx context.Context) ([]types.Position, error) {
	var positions []types.Position
	err := c.get(ctx, "/v1/portfolio", &positions)
	return positions, err
}

// GetPortfolio returns portfolio performance data.
func (c *Client) GetPortfolio(ctx context.Context) (types.Portfolio, error) {
	var portfolio types.Portfolio
	err := c.get(ctx, "/v1/portfolio/performance", &portfolio)
	return portfolio, err
}
