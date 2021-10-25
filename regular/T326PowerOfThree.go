package regular

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}

	for n > 1 {
		if n%3 != 0 {
			break
		}
		n /= 3
	}

	return n == 1
}
