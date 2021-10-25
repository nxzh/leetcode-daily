package regular

import "fmt"

func search(nums []int, target int) int {
	hi := len(nums) - 1
	lo := 0
	for lo <= hi {
		mid := (lo + hi) / 2
		if target > nums[mid] {
			lo = mid + 1
		} else if target < nums[mid] {
			hi = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func T704() {
	res := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Println(res)
}
