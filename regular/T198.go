package regular

import (
	"fmt"
	"math"
)

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = int(math.Max(float64(dp[0]), float64(nums[1])))
	for i := 2; i < len(nums); i++ {
		dp[i] = int(math.Max(float64(nums[i]+dp[i-2]), float64(dp[i-1])))
	}
	return dp[len(nums)-1]
}
func rob2(nums []int) int {
	dp := make(map[int]int)
	if len(nums) == 1 {
		return nums[0]
	}
	var internal func(i int) int
	internal = func(i int) int {
		if i < 0 {
			return 0
		}
		if i == 0 {
			return nums[0]
		}
		if _, presents := dp[i]; presents {
			return dp[i]
		}
		if _, presents := dp[i-2]; !presents {
			dp[i-2] = internal(i - 2)
		}
		if _, presents := dp[i-1]; !presents {
			dp[i-1] = internal(i - 1)
		}
		pick := nums[i] + dp[i-2]
		noPick := dp[i-1]
		val := int(math.Max(float64(pick), float64(noPick)))
		dp[i] = val
		return dp[i]
	}
	return internal(len(nums) - 1)
}

func rob3(nums []int) int {
	dp := make(map[int]int)
	var internal func(i int) int
	internal = func(i int) int {
		if i >= len(nums) {
			return 0
		}
		if val, exist := dp[i]; exist {
			return val
		}
		pick := nums[i] + internal(i+2)
		noPick := internal(i + 1)
		val := int(math.Max(float64(pick), float64(noPick)))
		dp[i] = val
		return val
	}
	return internal(0)
}

func T198() {
	count := rob([]int{2, 7, 9, 3, 1})
	fmt.Println(count)
}
