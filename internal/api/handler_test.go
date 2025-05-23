package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"orderpackscalculator/internal/config"
)

func init() {
	// Ensure default pack sizes are loaded for tests
	config.LoadDefaultPackSizes()
}

func TestCalculatePacksHandler_Success(t *testing.T) {
	body := CalculateRequest{
		PackSizes:   []int{250, 500, 1000},
		OrderAmount: 1250,
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/api/packs/calculate", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CalculatePacksHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", resp.StatusCode)
	}

	var out CalculateResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if out.Packs[1000] != 1 || out.Packs[250] != 1 {
		t.Errorf("unexpected packs: %v", out.Packs)
	}
}

func TestCalculatePacksHandler_WithOvershipping(t *testing.T) {
	body := CalculateRequest{
		PackSizes:   []int{250, 500},
		OrderAmount: 600, // Will fulfill with overshipping: 250 + 500 = 750
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/api/packs/calculate", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CalculatePacksHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", resp.StatusCode)
	}

	var out CalculateResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if out.Packs[250] != 1 || out.Packs[500] != 1 {
		t.Errorf("unexpected packs: %v", out.Packs)
	}
}

func TestCalculatePacksHandler_UsesDefaults(t *testing.T) {
	body := CalculateRequest{
		PackSizes:   []int{}, // Empty - should use defaults
		OrderAmount: 250,     // Should use one 250-pack with defaults
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/api/packs/calculate", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CalculatePacksHandler(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", resp.StatusCode)
	}

	var out CalculateResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	// Should use default pack sizes and return exactly one 250-pack
	if out.Packs[250] != 1 {
		t.Errorf("expected {250: 1}, got %v", out.Packs)
	}
}
