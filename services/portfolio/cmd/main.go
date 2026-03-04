package main

import (
	"log"
	"net"
	"os"

	"github.com/cnylum/accrue-engine/services/portfolio/internal/server"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9002"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("portfolio: failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	server.Register(srv)

	log.Printf("portfolio-service starting on :%s", port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("portfolio-service failed: %v", err)
	}
}
