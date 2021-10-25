package regular

import "math"

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	sq := int(math.Sqrt(float64(n)))
	cache := make(map[int]bool)
	for i := 2; i <= sq; i++ {
		_, prs := cache[i]
		if !prs {
			for j := i * i; j < n; j += i {
				cache[j] = false
			}
		}

	}
	count := 0
	for i := 2; i < n; i++ {
		_, prs := cache[i]
		if !prs {
			count++
		}
	}
	return count
}
