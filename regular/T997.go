package regular

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	posFirstNonNegative := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			posFirstNonNegative = i
			break
		}
	}
	left, right := posFirstNonNegative-1, posFirstNonNegative
	ret := make([]int, len(nums))
	pos := 0
	nextL := -1
	nextR := -1
	for pos < len(nums) {
		if nextL == -1 && left >= 0 {
			nextL = nums[left] * nums[left]
		}
		if nextR == -1 && right < len(nums) {
			nextR = nums[right] * nums[right]
		}
		if nextR < 0 || nextL >= 0 && nextL < nextR || nextL == nextR {
			ret[pos] = nextL
			nextL = -1
			left--
		} else if nextL < 0 || nextR >= 0 && nextR < nextL {
			ret[pos] = nextR
			nextR = -1
			right++
		}
		pos++
	}
	return ret
}

func T997() {
	res := sortedSquares([]int{-5, -3, -2, -1})
	fmt.Println(res)
}
