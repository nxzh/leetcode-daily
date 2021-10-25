package regular

func majorityElement(nums []int) int {
	m := make(map[int]int)
	size := len(nums)
	for i := 0; i < size; i++ {
		_, prs := m[nums[i]]
		if prs {
			m[nums[i]] = m[nums[i]] + 1
		} else {
			m[nums[i]] = 1

		}
		if m[nums[i]] > size/2 {
			return nums[i]
		}
	}
	return -1
}
