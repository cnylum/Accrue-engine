package position

// Position represents a held position in an asset.
type Position struct {
	AccountID     string
	Symbol        string
	AssetType     string
	Quantity      string
	CostBasis     string
	MarketValue   string
	UnrealizedPnL string
}

// Service manages positions.
type Service struct{}

// New creates a new position service.
func New() *Service {
	return &Service{}
}
