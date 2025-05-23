package api

import (
	"encoding/json"
	"net/http"

	"orderpackscalculator/internal/config"
	"orderpackscalculator/internal/packs"
)

// Request body for pack calculation
type CalculateRequest struct {
	PackSizes   []int `json:"pack_sizes"`
	OrderAmount int   `json:"order_amount"`
}

// Response body for pack calculation
type CalculateResponse struct {
	Packs map[int]int `json:"packs,omitempty"`
	Error string      `json:"error,omitempty"`
}

// CalculatePacksHandler handles POST /api/packs/calculate
func CalculatePacksHandler(w http.ResponseWriter, r *http.Request) {
	// CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight
	if r.Method == "OPTIONS" {
		return
	}

	var req CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(CalculateResponse{Error: "Invalid request body"})
		return
	}
	packSizes := req.PackSizes
	if len(packSizes) == 0 {
		packSizes = config.GetDefaultPackSizes()
	}
	result := packs.CalculatePacks(packSizes, req.OrderAmount)
	if result == nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(CalculateResponse{Error: "Cannot fulfill order exactly with given pack sizes."})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CalculateResponse{Packs: result})
}
