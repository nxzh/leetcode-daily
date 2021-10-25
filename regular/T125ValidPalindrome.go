package regular

import "strings"

func isPalindrome(s string) bool {
	lowS := strings.ToLower(s)
	l := 0
	r := len(s) - 1
	for l < r {
		if !isLetter(lowS[l]) {
			l++
			continue
		}
		if !isLetter(lowS[r]) {
			r--
			continue
		}

		if lowS[l] != lowS[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isLetter(r uint8) bool {
	if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
		return true
	}
	return false
}
