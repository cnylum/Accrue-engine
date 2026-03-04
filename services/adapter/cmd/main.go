package main

import (
	"log"
	"net"
	"os"

	"github.com/cnylum/accrue-engine/services/adapter/internal/server"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("adapter: failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	server.Register(srv)

	log.Printf("adapter-service starting on :%s", port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("adapter-service failed: %v", err)
	}
}
