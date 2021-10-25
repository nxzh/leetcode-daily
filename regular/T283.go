package regular

import "fmt"

func moveZeroes(nums []int) {
	// 1. set l,r
	// 2. l to find first zero
	// 3. r to find first non-zero number after l
	// 4. exchange the elements at l, r. l will be the non-zero number, r is zero
	// 5. l = l+1, r now is point to a zero-value element.
	// 6. if r+1 < len(nums), repeat 2-5

	l, r := len(nums), 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			l = i
			break
		}
	}
	r = l + 1
	for r < len(nums) {
		if nums[r] == 0 {
			r++
			continue
		}
		nums[l] = nums[r]
		nums[r] = 0
		l++
		r++
	}
}

func T283() {
	ary := []int{2, 1}
	moveZeroes(ary)
	fmt.Println(ary)
}
