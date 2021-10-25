package y2021m09

import "fmt"

func arrayNesting(nums []int) int {
	cache := make(map[int]int)
	max := 0
	for i := 0; i < len(nums); i++ {
		if i == nums[i] {
			cache[i] = 0
		} else {
			arrayNestingInternal(nums, nums[i], i, 1, cache)
		}
		if max < cache[i] {
			max = cache[i]
		}
	}
	return max
}
func arrayNestingInternal(nums []int, cur int, begin int, depth int, cache map[int]int) {
	if cur == begin {
		cache[begin] = depth
	} else {
		if _, prs := cache[cur]; !prs {
			arrayNestingInternal(nums, nums[cur], begin, depth+1, cache)
			cache[cur] = depth
		}
	}
}
func Call01() {
	res := arrayNesting([]int{5, 4, 0, 3, 1, 6, 2})
	fmt.Println(res)
}
