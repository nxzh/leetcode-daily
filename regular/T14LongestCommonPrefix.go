package regular

func longestCommonPrefix(strs []string) string {
	strCnt := len(strs)
	result := ""
	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i]
		count := 0
		for j := 0; j < len(strs); j++ {
			if i == len(strs[j]) {
				break
			}
			if ch == strs[j][i] {
				count++
			}
		}
		if count != strCnt {
			break
		} else {
			result += string(ch)
		}
	}
	return result
}
