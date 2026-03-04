// Package types defines domain types shared across the SDK.
package types

// Order represents a trade order.
type Order struct {
	ID         string `json:"id"`
	AccountID  string `json:"account_id"`
	Symbol     string `json:"symbol"`
	Side       string `json:"side"`
	OrderType  string `json:"order_type"`
	Quantity   string `json:"quantity"`
	LimitPrice string `json:"limit_price,omitempty"`
	FilledQty  string `json:"filled_quantity"`
	Status     string `json:"status"`
}

// Position represents a held position.
type Position struct {
	Symbol        string `json:"symbol"`
	AssetType     string `json:"asset_type"`
	Quantity      string `json:"quantity"`
	CostBasis     string `json:"cost_basis"`
	MarketValue   string `json:"market_value"`
	UnrealizedPnL string `json:"unrealized_pnl"`
}

// Balance represents an account balance.
type Balance struct {
	Available string `json:"available"`
	Held      string `json:"held"`
	Total     string `json:"total"`
	Currency  string `json:"currency"`
}

// Portfolio holds portfolio performance data.
type Portfolio struct {
	Positions  []Position `json:"positions"`
	TotalValue string     `json:"total_value"`
	TotalCost  string     `json:"total_cost"`
	TotalPnL   string     `json:"total_pnl"`
	ReturnPct  string     `json:"return_pct"`
}

// Account holds basic account information.
type Account struct {
	ID      string  `json:"id"`
	Balance Balance `json:"balance"`
}

// PlaceOrderRequest is the input for placing an order.
type PlaceOrderRequest struct {
	Symbol     string `json:"symbol"`
	Side       string `json:"side"`
	OrderType  string `json:"order_type"`
	Quantity   string `json:"quantity"`
	LimitPrice string `json:"limit_price,omitempty"`
}
