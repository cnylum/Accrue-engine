package server

import "google.golang.org/grpc"

// Register registers the portfolio service with the gRPC server.
func Register(srv *grpc.Server) {
	// Will register portfoliov1.PortfolioServiceServer once proto codegen is wired.
}
