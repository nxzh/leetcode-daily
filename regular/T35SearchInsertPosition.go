package regular

import "fmt"

func searchInsert(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1
	mid := 0
	for lo <= hi {
		mid = (lo + hi) / 2
		if target > nums[mid] {
			lo = mid + 1
		} else if target < nums[mid] {
			hi = mid - 1
		} else {
			return mid
		}
	}
	return lo
}

func T35() {
	res := searchInsert([]int{1, 2, 4, 6, 7}, 3)
	fmt.Println(res)
}
