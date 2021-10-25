package regular

import "strings"

func findDifferentBinaryString(nums []string) string {
	cache := make(map[string]bool, len(nums))
	for _, num := range nums {
		cache[num] = true
	}
	initial := strings.Repeat("0", len(nums))
	return findInternally(initial, cache)
}

func findInternally(cur string, cache map[string]bool) string {
	_, prs := cache[cur]
	if !prs {
		return cur
	}
	lenStr := len(cur)
	if cur[lenStr-1] == '0' {
		next := cur[:lenStr-1] + "1"
		return findInternally(next, cache)
	} else {
		i0 := strings.LastIndex(cur, "0")
		if i0 == -1 {
			return ""
		}
		next := cur[:i0] + "1"
		for i := 0; i < lenStr-i0-1; i++ {
			next += "0"
		}
		return findInternally(next, cache)
	}
}
