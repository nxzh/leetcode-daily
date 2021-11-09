package regular

import "fmt"

func kthDistinct(arr []string, k int) string {
	cache := make(map[string]int)
	for _, s := range arr {
		cache[s] += 1
	}
	c := 1
	for _, s := range arr {
		if cache[s] == 1 {
			if c == k {
				return s
			} else {
				c++
			}
		}
	}
	return ""
}

func T5898() {
	s := kthDistinct([]string{"aaa", "aa", "a"}, 1)
	fmt.Println(s)
}
