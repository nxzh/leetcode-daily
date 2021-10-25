package regular

import "fmt"

func construct2DArray(original []int, m int, n int) [][]int {
	// r, c
	// if r * c != m * n return original
	// for i from r*c, res[i/n][i%m] = o[i][j]

	count := len(original)
	if count != m*n {
		return [][]int{}
	}
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < count; i++ {
		res[i/n][i%n] = original[i]
	}
	return res
}

func T5817() {
	res := construct2DArray([]int{3}, 1, 2)
	fmt.Println(res)
}
