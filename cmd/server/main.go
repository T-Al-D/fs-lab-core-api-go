package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/t-al-d/fs-lab-core-api-go/internal/router"
)

// main is the entry point of the application.
func main() {
	// Port makes the server more adapable
		port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback for lokal
	}

	addr := ":" + port
	fmt.Println("Starting HTTP server on", addr)

	http.ListenAndServe(addr, router.New())
}