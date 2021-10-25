package regular

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[l] {
			l += 1
			if i != l {
				nums[l] = nums[i]
			}
		}
	}
	return l + 1
}
