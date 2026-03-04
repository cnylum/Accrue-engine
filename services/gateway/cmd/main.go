package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cnylum/accrue-engine/services/gateway/internal/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := router.New()

	log.Printf("gateway starting on :%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("gateway failed: %v", err)
	}
}
