package y2021m08

import "math"

func stoneGame(piles []int) bool {
	totalStones := 0
	for _, stones := range piles {
		totalStones += stones
	}
	target := totalStones/2 + 1
	cache := make([][]int, len(piles))
	for i := 0; i < len(piles); i++ {
		cache[i] = make([]int, len(piles))
	}
	return pick(piles, 0, len(piles)-1, target, 0, cache)
}

func pick(piles []int, l int, r int, target int, current int, cache [][]int) bool {
	if cache[l][r-1] > current {
		return false
	} else {
		cache[l][r-1] = current
	}
	if l >= r {
		return false
	}
	head := piles[0]
	tail := piles[r-1]
	if current+head >= target || current+tail >= target {
		return true
	} else if r-l == 0 {
		return false
	} else {
		return pick(piles, 2, r, target, current+head, cache) ||
			pick(piles, 1, r-1, target, int(math.Max(float64(current+head), float64(current+tail))), cache) ||
			pick(piles, 0, r-2, target, current+tail, cache)
	}
}
