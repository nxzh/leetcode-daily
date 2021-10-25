package y2021m08

import "math"

func minWindow(s string, t string) string {
	tCount, sCount := make([]int, 52), make([]int, 52)
	for i := 0; i < len(t); i++ {
		tCount[charToIndex(t[i])] = tCount[charToIndex(t[i])] + 1
	}
	minLen, minR, l, r := math.MaxInt32, 0, 0, 0
	for r < len(s) {
		sCount[charToIndex(s[r])] = sCount[charToIndex(s[r])] + 1
		for hasTInWindow(&tCount, &sCount) {
			if r-l+1 < minLen {
				minR = r
				minLen = r - l + 1
			}
			sCount[charToIndex(s[l])] = sCount[charToIndex(s[l])] - 1
			l++
		}
		r++
	}
	if minLen != math.MaxInt32 {
		return s[minR-minLen+1 : minR+1]
	}
	return ""
}

func hasTInWindow(tCount *[]int, sCount *[]int) bool {
	for i := 0; i < len(*tCount); i++ {
		if (*sCount)[i] < (*tCount)[i] {
			return false
		}
	}
	return true
}

func charToIndex(ch byte) int {
	if ch >= 'A' && ch <= 'Z' {
		return int(ch - 'A' + 26)
	} else {
		return int(ch - 'a')
	}
}
