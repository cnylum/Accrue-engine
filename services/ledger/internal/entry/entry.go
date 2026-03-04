package entry

// Entry represents a single ledger entry (one side of a double-entry).
type Entry struct {
	ID        string
	AccountID string
	Amount    string
	Currency  string
	Direction string // "credit" or "debit"
	Reference string
}

// Service manages ledger entries.
type Service struct{}

// New creates a new entry service.
func New() *Service {
	return &Service{}
}
