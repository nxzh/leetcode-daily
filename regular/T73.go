package regular

func setZeroes(matrix [][]int) {
	rows := make([]int, len(matrix))
	cols := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = 1
				cols[j] = 1
			}
		}
	}
	for i := 0; i < len(rows); i++ {
		if rows[i] == 1 {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 0; i < len(cols); i++ {
		if cols[i] == 1 {
			for j := 0; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}
}
