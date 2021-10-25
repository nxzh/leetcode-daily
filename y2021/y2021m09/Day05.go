package y2021m09

import (
	"fmt"
	"sort"
	"strings"
)

func orderlyQueue(s string, k int) string {
	hist := make(map[string]bool)
	sb := make([]byte, k)
	for i := 0; i < k; i++ {
		sb[i] = s[i]
	}
	sort.Slice(sb, func(i, j int) bool {
		return sb[i] < sb[j]
	})

	for i := k; i < len(s); i++ {
		sb = append(sb, s[i])
	}
	hist[s] = true
	min := s
	for {
		before := string(sb)
		fmt.Println(before)
		max := sb[k-1]
		for i := k; i < len(sb); i++ {
			sb[i-1] = sb[i]
		}
		sb[len(sb)-1] = max
		for i := k - 1; i > 0; i-- {
			if sb[i] < sb[i-1] {
				temp := sb[i-1]
				sb[i-1] = sb[i]
				sb[i] = temp
			}
		}
		cur := string(sb)
		if strings.Compare(cur, min) < 0 {
			min = cur
		}
		if hist[cur] {
			return min
		}
		hist[cur] = true
	}
}

func Call05() {
	res := orderlyQueue("xmvzi", 2)
	fmt.Println(res)
}
