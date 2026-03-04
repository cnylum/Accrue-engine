package handler

// Handler holds dependencies for HTTP handlers.
type Handler struct{}

// New creates a new Handler.
func New() *Handler {
	return &Handler{}
}
