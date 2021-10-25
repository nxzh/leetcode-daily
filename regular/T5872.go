package regular

import (
	"fmt"
	"sort"
	"strings"
)

func numOfPairs(nums []string, target string) int {
	// sort the string
	// until target.indexof(num[i]) < 0
	// 	toFind = target.subStr(len(num[i])) , binary search from the remaining, if found, count ++

	sort.Strings(nums)
	count := 0
	n := len(nums)
	findRemainFromRight := func(index int, toFind string) int {
		l, h := 0, n-1
		for l <= h {
			mid := (l + h) / 2
			midStr := nums[mid]
			comp := strings.Compare(toFind, midStr)
			if comp > 0 {
				l = mid + 1
			} else if comp < 0 {
				h = mid - 1
			} else {
				ret := 0
				for i := mid; i < n; i++ {
					if strings.Compare(toFind, nums[i]) == 0 {
						if index != i {
							ret++
						}
					} else {
						break
					}
				}
				for i := mid - 1; i >= 0; i-- {
					if strings.Compare(toFind, nums[i]) == 0 {
						if index != i {
							ret++
						}
					} else {
						break
					}
				}
				return ret
			}
		}
		return 0
	}
	for i := 0; i < n; i++ {
		if len(nums[i]) >= len(target) {
			continue
		}
		if strings.Compare(nums[i], target[0:len(nums[i])]) == 0 {
			remain := target[len(nums[i]):]
			ret := findRemainFromRight(i, remain)
			count += ret
		}
	}
	return count
}
func T5872() {
	res := numOfPairs([]string{"1", "111"}, "11")
	fmt.Println(res)
}
