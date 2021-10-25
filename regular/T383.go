package regular

func canConstruct(ransomNote string, magazine string) bool {
	m := len(ransomNote)
	n := len(magazine)
	if m > n {
		return false
	}
	var cache1 [26]int
	var cache2 [26]int
	for i := 0; i < m; i++ {
		cache1[ransomNote[i]-'a']++
	}
	for i := 0; i < n; i++ {
		cache2[magazine[i]-'a']++
	}
	for i := 0; i < len(cache1); i++ {
		if cache1[i] > cache2[i] {
			return false
		}
	}
	return true
}
