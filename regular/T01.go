package regular

import (
	"fmt"
)

func twoSum1(nums []int, target int) []int {
	cache := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, prs := cache[target-nums[i]]; prs {
			return []int{v, i}
		}
		cache[nums[i]] = i
	}
	return nil
}

func T01() {
	res := twoSum1([]int{0, 1, 2}, 1)
	fmt.Println(res)
}
