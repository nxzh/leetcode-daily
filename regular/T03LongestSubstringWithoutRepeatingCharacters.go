package regular

import (
	"fmt"
	"math"
)

func lengthOfLongestSubstring(s string) int {
	// create a map to record the appearance times
	// set l, r := 0, 0
	// r iterate from 0 to len(s), map[e]++
	// if map[e] == 2, that means it has duplicated
	//     then curLen = r-l , compare and set max
	//     then l++, and remove the elements from map at l, until map[e] is 0
	//     l is porint to the first e+1
	// else increase curLen

	l, r := 0, 0
	c := make([]byte, 256)
	max := 0
	for r < len(s) {
		cur := s[r]
		c[cur]++
		if c[cur] == 2 {
			max = int(math.Max(float64(max), float64(r-l)))
			for c[cur] != 1 {
				rm := s[l]
				c[rm]--
				l++
			}
		}
		r++
	}
	max = int(math.Max(float64(max), float64(r-l)))
	return max
}
func CallT03() {
	res := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(res)
}
