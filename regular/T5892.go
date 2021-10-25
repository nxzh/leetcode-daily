package regular

func stoneGameIX(stones []int) bool {
	// sort
	// remove all the 3x
	// if for
	var n [3]int
	for i := 0; i < len(stones); i++ {
		n[stones[i]%3]++
	}
	n[0] = n[0] % 2
	n[1] = n[1] % 2
	n[2] = n[2] % 2
	sum := n[0] + n[1] + n[2]
	if sum != 3 {
		return false
	}

	return true
}
