package pnl

// Performance holds portfolio performance metrics.
type Performance struct {
	TotalValue string
	TotalCost  string
	TotalPnL   string
	ReturnPct  string
}

// Calculator computes P&L metrics.
type Calculator struct{}

// New creates a new P&L calculator.
func New() *Calculator {
	return &Calculator{}
}
