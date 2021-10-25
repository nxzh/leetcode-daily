package regular

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	l := 0
	r := len(needle) - 1

	for r < len(haystack) {
		if haystack[l:r+1] == needle {
			return l
		}
		l++
		r++
	}
	return -1
}
