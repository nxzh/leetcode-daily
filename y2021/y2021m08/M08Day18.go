package y2021m08

func numDecodings(s string) int {
	cache := make(map[string]int)
	return decodings(s, &cache)
}

func decodings(s string, cache *map[string]int) int {
	if len(s) == 0 {
		return 1
	}
	if len(s) == 1 {
		if s[0] == '0' {
			return 0
		}
		return 1
	} else {
		if s[0] == '0' {
			return 0
		}
		str1 := s[1:]
		_, prs := (*cache)[str1]
		if !prs {
			(*cache)[str1] = decodings(str1, cache)
		}
		count1 := (*cache)[str1]
		if s[0] >= '2' && s[1] >= '7' || s[0] > '2' {
			return count1
		}
		str2 := s[2:]
		_, prs = (*cache)[str2]
		if !prs {
			(*cache)[str2] = decodings(s[2:], cache)
		}
		count2 := (*cache)[str2]
		return count1 + count2
	}
}
