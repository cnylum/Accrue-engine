package main

import (
	"log"
	"net"
	"os"

	"github.com/cnylum/accrue-engine/services/ledger/internal/server"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9003"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("ledger: failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	server.Register(srv)

	log.Printf("ledger-service starting on :%s", port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("ledger-service failed: %v", err)
	}
}
