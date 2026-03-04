package server

import "google.golang.org/grpc"

// Register registers the ledger service with the gRPC server.
func Register(srv *grpc.Server) {
	// Will register ledgerv1.LedgerServiceServer once proto codegen is wired.
}
