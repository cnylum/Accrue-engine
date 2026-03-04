package store

// Store is the PostgreSQL storage layer for portfolio data.
type Store struct{}

// New creates a new store.
func New() *Store {
	return &Store{}
}
