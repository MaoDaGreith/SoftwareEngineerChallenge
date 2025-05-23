package config

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

var defaultPackSizes []int

// LoadDefaultPackSizes loads pack sizes from env var, config file, or defaults
func LoadDefaultPackSizes() {
	// Try environment variable first
	if env := os.Getenv("PACK_SIZES"); env != "" {
		parts := strings.Split(env, ",")
		var sizes []int
		for _, p := range parts {
			if n, err := strconv.Atoi(strings.TrimSpace(p)); err == nil {
				sizes = append(sizes, n)
			}
		}
		if len(sizes) > 0 {
			defaultPackSizes = sizes
			return
		}
	}
	// Try config.json
	f, err := os.Open("config.json")
	if err == nil {
		defer f.Close()
		var cfg struct {
			PackSizes []int `json:"pack_sizes"`
		}
		if err := json.NewDecoder(f).Decode(&cfg); err == nil && len(cfg.PackSizes) > 0 {
			defaultPackSizes = cfg.PackSizes
			return
		}
	}
	// Default fallback
	defaultPackSizes = []int{250, 500, 1000, 2000, 5000}
}

// GetDefaultPackSizes returns the loaded pack sizes
func GetDefaultPackSizes() []int {
	return defaultPackSizes
}
