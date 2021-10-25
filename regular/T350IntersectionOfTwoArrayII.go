package regular

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	var m = make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] += 1
	}

	for i := 0; i < len(nums2); i++ {
		key := nums2[i]
		_, prs := m[key]
		if prs {
			res = append(res, key)
			m[key] -= 1
			if m[key] == 0 {
				delete(m, key)
			}
		}
	}
	return res
}
