package y2021m08

import "fmt"

func findMin(nums []int) int {
	high := len(nums) - 1
	low := 0
	for nums[low] > nums[high] {
		mid := (low + high) / 2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

func Call0831() {
	//res := findMin([]int{4, 5, 6, 7, 0, 1, 2})
	//res := findMinin([]int{4, 5, 6, 7, 0, 1, 2})
	res := findMinin([]int{7, 1, 2, 3, 4, 5, 6})
	fmt.Println(res)
}
