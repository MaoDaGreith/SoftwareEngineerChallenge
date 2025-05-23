package packs

import (
	"reflect"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	tests := []struct {
		name      string
		packSizes []int
		order     int
		expect    map[int]int
	}{
		{
			name:      "Single pack size, exact",
			packSizes: []int{250},
			order:     500,
			expect:    map[int]int{250: 2},
		},
		{
			name:      "Multiple pack sizes, optimal",
			packSizes: []int{250, 500, 1000},
			order:     1250,
			expect:    map[int]int{1000: 1, 250: 1},
		},
		{
			name:      "Cannot fulfill exactly - minimal overshipping",
			packSizes: []int{250, 500},
			order:     600,
			expect:    map[int]int{250: 1, 500: 1}, // 750 total, overship by 150, fewer packs than 3x250
		},
		{
			name:      "Edge case: zero order",
			packSizes: []int{250, 500},
			order:     0,
			expect:    nil,
		},
		{
			name:      "Edge case: empty pack sizes",
			packSizes: []int{},
			order:     100,
			expect:    nil,
		},
		{
			name:      "Large order, edge pack sizes",
			packSizes: []int{23, 31, 53},
			order:     500000,
			expect:    map[int]int{23: 2, 31: 7, 53: 9429},
		},
		{
			name:      "Prefer fewer packs with same overshipping",
			packSizes: []int{500, 1000},
			order:     780,
			expect:    map[int]int{1000: 1}, // 1x1000 is better than 2x500
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CalculatePacks(tc.packSizes, tc.order)
			if !reflect.DeepEqual(got, tc.expect) {
				t.Errorf("got %v, want %v", got, tc.expect)
			}
		})
	}
}
