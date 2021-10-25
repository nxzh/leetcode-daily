package regular

func plusOne(digits []int) []int {
	s := make([]int, len(digits)+1)
	carry := 1
	i := len(digits) - 1
	for carry > 0 && i >= 0 {
		s[i+1] = (digits[i] + carry) % 10
		carry = (digits[i] + carry) / 10
		i--
	}
	if carry > 0 {
		s[0] = 1
		return s
	} else {
		for j := i; j >= 0; j-- {
			s[j+1] = digits[j]
		}
		return s[1:]
	}
}
