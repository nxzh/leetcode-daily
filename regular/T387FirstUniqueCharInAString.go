package regular

func firstUniqChar(s string) int {
	chars := make([]int, 26)
	for i := 0; i < len(s); i++ {
		chars[s[i]-'a'] += 1
	}

	for i := 0; i < len(s); i++ {
		if chars[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
