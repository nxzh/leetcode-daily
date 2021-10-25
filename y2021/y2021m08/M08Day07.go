package y2021m08

import (
	"math"
)

func minCut(s string) int {
	cache := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		cache[i] = make([]int, len(s))
		for j := 0; j < len(s); j++ {
			cache[i][j] = -1
		}
	}
	return solution0807(s, 0, len(s)-1, 0, &cache)
}
func solution0807(s string, l int, r int, level int, cache *[][]int) int {
	if l >= r {
		return level
	}
	if isPalindrome0807(s, l, r) {
		res := 0
		if l > 0 {
			res += 1 + solution0807(s[0:l], 0, l-1, 0, cache)
		}
		if r < len(s)-1 {
			res += 1 + solution0807(s[r+1:], 0, len(s)-r-1-1, 0, cache)
		}
		return res
	} else {
		res := int(math.Min(float64(solution0807(s, l+1, r, level+1, cache)), float64(solution0807(s, l, r-1, level+1, cache))))
		return res
	}
}
func isPalindrome0807(s string, l int, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
