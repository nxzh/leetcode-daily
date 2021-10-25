package y2021m08

func groupAnagrams(strs []string) [][]string {
	index := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		b := make([]byte, 26)
		for j := 0; j < len(strs[i]); j++ {
			b[strs[i][j]-'a']++
		}
		index[string(b)] = append(index[string(b)], strs[i])
	}
	var ret [][]string
	for _, strings := range index {
		ret = append(ret, strings)
	}
	return ret
}
