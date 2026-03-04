package router

import "net/http"

func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})
	// Placeholder routes — will be wired to handlers
	mux.HandleFunc("GET /v1/orders", notImplemented)
	mux.HandleFunc("POST /v1/orders", notImplemented)
	mux.HandleFunc("GET /v1/portfolio", notImplemented)
	mux.HandleFunc("GET /v1/account/balance", notImplemented)
	return mux
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(`{"error":"not implemented"}`))
}
