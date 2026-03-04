package client

import (
	"context"

	"github.com/cnylum/accrue-engine/sdk/types"
)

// GetBalance returns the account balance.
func (c *Client) GetBalance(ctx context.Context) (types.Balance, error) {
	var balance types.Balance
	err := c.get(ctx, "/v1/account/balance", &balance)
	return balance, err
}

// GetAccount returns account information.
func (c *Client) GetAccount(ctx context.Context) (types.Account, error) {
	var account types.Account
	err := c.get(ctx, "/v1/account", &account)
	return account, err
}
