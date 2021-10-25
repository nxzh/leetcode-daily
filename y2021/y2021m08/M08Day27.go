package y2021m08

import (
	"fmt"
	"sort"
	"strings"
)

func findLUSlength(strs []string) int {
	counter := make(map[string]int)
	sort.Slice(strs, func(i, j int) bool {
		if len(strs[i]) == len(strs[j]) {
			return strings.Compare(strs[i], strs[j]) < 0
		} else {
			return len(strs[i]) < len(strs[j])
		}
	})
	strsL := len(strs)
	uniqueStrs := []string{strs[0]}
	counter[strs[0]]++
	for i := 1; i < strsL; i++ {
		counter[strs[i]] += 1
		if strings.Compare(strs[i], uniqueStrs[len(uniqueStrs)-1]) != 0 {
			uniqueStrs = append(uniqueStrs, strs[i])
		}
	}
	unqStrsL := len(uniqueStrs)
	if counter[uniqueStrs[unqStrsL-1]] == 1 {
		return len(uniqueStrs[unqStrsL-1])
	}
	for i := len(strs) - 1; i >= 0; i-- {
		cur := strs[i]
		if counter[cur] == 1 {
			found := false
			for j := i + 1; !found && j < len(strs); j++ {
				if isSubstr(strs[j], cur) {
					found = true
				}
			}
			if !found {
				return len(cur)
			}
		}
	}
	return -1

}

func isSubstr(longer string, shorter string) bool {
	allowDiff := len(longer) - len(shorter)
	diff := 0
	for i, j := 0, 0; diff <= allowDiff && i < len(longer) && j < len(shorter); i++ {
		if longer[i] != shorter[j] {
			diff++
		} else {
			j++
		}
	}
	return diff <= allowDiff
}
func Call0827() {
	res := findLUSlength([]string{"abcabc", "abcabc", "abcabc", "abc", "abc", "aab"})
	fmt.Println(res)
}
