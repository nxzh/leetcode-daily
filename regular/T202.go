package regular

func isHappy(n int) bool {
	m := make(map[int]int)
	m[n] = 1
	for {
		res := 0
		for n > 0 {
			k := n % 10
			res += k * k
			n /= 10
		}
		if res == 1 {
			return true
		} else {
			_, prs := m[res]
			if prs {
				return false
			} else {
				m[res] = 1
				n = res
			}
		}
	}
}
