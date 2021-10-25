package y2021m08

import "sort"

func twoSum(nums []int, target int) []int {
	ary := make([]int, len(nums))
	copy(ary, nums)
	sort.Ints(ary)

	l, r := 0, len(ary)-1
	var res []int
	for l < r {
		sum := ary[l] + ary[r]
		if sum == target {
			res = []int{ary[l], ary[r]}
			break
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	var resIndex []int
	for i := 0; i < len(nums); i++ {
		if res[0] == nums[i] {
			resIndex = append(resIndex, i)
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if res[1] == nums[i] {
			resIndex = append(resIndex, i)
			break
		}
	}
	return resIndex
}
