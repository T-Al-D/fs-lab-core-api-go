package router

import (
	"net/http"

	"github.com/t-al-d/fs-lab-core-api-go/internal/handlers"
)

// New returns the HTTP router
func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/meta", handlers.Meta)

	return mux
}
