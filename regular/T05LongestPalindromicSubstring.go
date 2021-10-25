package regular

import "fmt"

func CallT05() {
	res := longestPalindrome("babaddtattarrattatddetartrateedredividerb")
	fmt.Println(res)
}
func longestPalindrome(s string) string {
	cache := make(map[string]string)
	return solutionT05(s, cache)
}

func solutionT05(s string, cache map[string]string) string {
	if val, prs := cache[s]; prs {
		return val
	}
	if isPalindromic(s) {
		cache[s] = s
		return s
	}
	str1 := s[1:]
	str2 := s[:len(s)-1]
	if _, prs := cache[str1]; !prs {
		cache[str1] = longestPalindrome(str1)
	}
	if _, prs := cache[str2]; !prs {
		cache[str2] = longestPalindrome(str2)
	}
	s1 := cache[str1]
	s2 := cache[str2]
	if len(s1) > len(s2) {
		return s1
	} else {
		return s2
	}
}

func isPalindromic(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
