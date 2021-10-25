package regular

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	// 1. set cache1, cache2 to count the current occurrence in len(s1)
	if len(s1) > len(s2) {
		return false
	}
	var cache1 [26]byte
	var cache2 [26]byte
	for i, ch := range s1 {
		cache1[ch-'a']++
		cache2[s2[i]-'a']++
	}
	if cache2 == cache1 {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		cache2[s2[i]-'a']++
		cache2[s2[i-len(s1)]-'a']--
		if cache1 == cache2 {
			return true
		}
	}

	return false
}

func T567() {
	res := checkInclusion("trinitrophenylmethylnitramine", "dinitrophenylhydrazinetrinitrophenylmethylnitramine")
	fmt.Println(res)
}
