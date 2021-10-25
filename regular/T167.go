package regular

import "fmt"

func twoSum(numbers []int, target int) []int {
	// 1. l, r -> 0, len-1
	// 2. if (nums[l] + nums[r] > target) {
	//     r--
	// } else if (nums[l] + nums[r] < target) {
	//     l++
	// } else add to

	l, r := 0, len(numbers)-1
	res := []int{}
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			res = append(res, l, r)
			l++
			r--
		}
	}
	return res
}

func T167() {
	res := twoSum([]int{}, 0)
	fmt.Println(res)
}
