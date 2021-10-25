package regular

func climbStairs(n int) int {
	m := make(map[int]int)
	return climb(n, m)
}

func climb(n int, m map[int]int) int {
	if n <= 1 {
		return 1
	}
	valOne, prs := m[n-1]
	if !prs {
		valOne = climb(n-1, m)
		m[n-1] = valOne
	}
	valTwo, prs := m[n-2]
	if !prs {
		valTwo = climb(n-2, m)
		m[n-2] = valTwo
	}
	return valOne + valTwo
}
