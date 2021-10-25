package regular

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	i1, i2 := 0, 0
	for i1 < m {
		if nums1[i1] > nums2[i2] {
			nums1[i1], nums2[i2] = nums2[i2], nums1[i1]
			for i := 0; i < n-1; i++ {
				if nums2[i] > nums2[i+1] {
					nums2[i], nums2[i+1] = nums2[i+1], nums2[i]
				}
			}
		}
		i1++
	}
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}
}

func T88() {
	ary1 := []int{1, 2, 3, 0, 0, 0}
	ary2 := []int{2, 5, 6}
	merge(ary1, 3, ary2, 3)
	fmt.Println(ary1)
}
