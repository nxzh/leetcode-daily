package regular

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, prs := m[nums[i]]
		if prs {
			delete(m, nums[i])
		} else {
			m[nums[i]] = 1
		}
	}
	for k := range m {
		return k
	}
	return 0
}
