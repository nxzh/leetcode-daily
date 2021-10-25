package y2021m08

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	minH := m
	minV := n

	for _, op := range ops {
		if op[0] < minH {
			minH = op[0]
		}
		if op[1] < minV {
			minV = op[1]
		}
	}
	return minV * minH
}
