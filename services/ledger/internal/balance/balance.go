package balance

// Balance holds account balance information.
type Balance struct {
	AccountID string
	Available string
	Held      string
	Total     string
	Currency  string
}

// Service manages balance calculations.
type Service struct{}

// New creates a new balance service.
func New() *Service {
	return &Service{}
}
