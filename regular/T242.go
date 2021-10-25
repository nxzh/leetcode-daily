package regular

func isAnagram(s string, t string) bool {
	m := len(s)
	n := len(t)
	if m != n {
		return false
	}
	var cacheS [26]int
	var cacheT [26]int

	for i := 0; i < m; i++ {
		cacheS[s[i]-'a']++
		cacheT[s[i]-'a']++
	}
	return cacheS == cacheT
}
