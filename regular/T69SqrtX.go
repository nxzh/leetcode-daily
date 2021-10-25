package regular

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	digits := 0
	t := x
	for t > 0 {
		digits++
		t = t / 10
	}
	digits = (digits + 1) / 2
	max := 1
	for digits > 0 {
		max *= 10
		digits--
	}
	min := max / 10
	for true {
		mid := (min + max) / 2
		tLeft := mid * mid
		vRight := (mid + 1) * (mid + 1)
		if tLeft <= x && vRight > x {
			return mid
		} else if vRight <= x {
			min = mid + 1
		} else if tLeft > x {
			max = mid
		}
	}
	return -1
}
