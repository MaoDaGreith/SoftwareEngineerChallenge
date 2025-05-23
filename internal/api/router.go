package api

import (
	"net/http"
)

// NewRouter sets up HTTP routes
func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("/api/packs/calculate", CalculatePacksHandler)

	// Static files
	fileServer := http.FileServer(http.Dir("web/"))
	mux.Handle("/", fileServer)

	return mux
}
