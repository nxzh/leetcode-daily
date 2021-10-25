package regular

func missingNumber(nums []int) int {
	length := len(nums)
	sumExpected := length * (length + 1) / 2
	sum := 0
	for i := 0; i < length; i++ {
		sum += nums[i]
	}
	return sumExpected - sum
}
