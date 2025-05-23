package packs

import "sort"

// CalculatePacks finds the optimal pack combination for an order
func CalculatePacks(packSizes []int, orderAmount int) map[int]int {
	if orderAmount <= 0 || len(packSizes) == 0 {
		return nil
	}

	sizes := make([]int, len(packSizes))
	copy(sizes, packSizes)
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	type state struct {
		packs     int
		breakdown map[int]int
		lastPack  int
	}

	maxPack := sizes[0]
	maxTotal := orderAmount + maxPack
	dp := make([]*state, maxTotal+1)
	dp[0] = &state{packs: 0, breakdown: map[int]int{}, lastPack: 0}

	for i := 1; i <= maxTotal; i++ {
		for _, sz := range sizes {
			if i >= sz && dp[i-sz] != nil {
				prev := dp[i-sz]
				newPacks := prev.packs + 1
				newBreakdown := make(map[int]int, len(prev.breakdown))
				for k, v := range prev.breakdown {
					newBreakdown[k] = v
				}
				newBreakdown[sz]++
				if dp[i] == nil || newPacks < dp[i].packs || (newPacks == dp[i].packs && sz > dp[i].lastPack) {
					dp[i] = &state{packs: newPacks, breakdown: newBreakdown, lastPack: sz}
				}
			}
		}
	}

	minOver := -1
	minPacks := -1
	var best map[int]int
	for total := orderAmount; total <= maxTotal; total++ {
		if dp[total] != nil {
			over := total - orderAmount
			if minOver == -1 || over < minOver || (over == minOver && (minPacks == -1 || dp[total].packs < minPacks)) {
				minOver = over
				minPacks = dp[total].packs
				best = dp[total].breakdown
			}
		}
	}
	if best == nil {
		return nil
	}
	return best
}
