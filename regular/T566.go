package regular

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	num := len(mat) * len(mat[0])
	cols := num / r
	if cols != c || r*c != num {
		return mat
	}
	reshaped := make([][]int, r)
	for i := 0; i < len(reshaped); i++ {
		reshaped[i] = make([]int, cols)
	}
	matC := len(mat[0])

	filled := 0
	for filled < num {
		reshaped[filled/c][filled%c] = mat[filled/matC][filled%matC]
		filled++
	}
	return reshaped
}

func T566() {
	res := matrixReshape([][]int{{1, 2}, {3, 4}}, 2, 4)
	fmt.Println(res)
}
