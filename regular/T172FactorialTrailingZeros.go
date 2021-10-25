package regular

func trailingZeroes(n int) int {
	count := 0
	if n/5 < 1 {
		return 0
	}
	count += n/5 + n/25 + n/125 + n/625 + n/3125
	return count
}
