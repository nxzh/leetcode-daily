package regular

import "fmt"

func rotate(nums []int, k int) {
	if len(nums) == 1 {
		return
	}
	dist := k % len(nums)
	if dist == 0 {
		return
	}
	i := 0
	next := nums[i]
	scanned := 0
	count := 0
	for count < len(nums) {
		pos := (i + dist) % len(nums)

		temp := nums[pos]
		nums[pos] = next
		if pos <= scanned {
			i = pos + 1
			next = nums[i]
			scanned++
		} else {
			i = pos
			next = temp
		}
		count++
	}
}
func T189() {
	ary := []int{-1, -100, 3, 99}
	rotate(ary, 2)
	fmt.Println(ary)
}
