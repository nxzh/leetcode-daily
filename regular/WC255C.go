package regular

import (
	"math"
	"sort"
)

func minimizeTheDifference(mat [][]int, target int) int {
	M := len(mat)
	N := len(mat[0])

	minVal := 0
	maxVal := 0
	for i := 0; i < M; i++ {
		sort.Ints(mat[i])
		minVal += mat[i][0]
		maxVal += mat[i][N-1]
		for j := N - 1; j > 0; j-- {
			mat[i][j] = mat[i][j] - mat[i][j-1]
		}
	}
	minDiff := math.MaxInt32
	if target <= minVal {
		return minVal - target
	} else if target >= maxVal {
		return target - maxVal
	} else {
		base := minVal
		for i := 0; i < M; i++ {
			for j := 1; j < N; j++ {
				base += mat[i][j]
				diff := int(math.Abs(float64(target - base)))
				if diff < minDiff {
					minDiff = diff
				}
			}
		}
	}
	return minDiff
}

func findFromMatrix(base, M int, N int, target int) {
	for i := 0; i < M; i++ {
		for j := 1; j < N; j++ {

		}
	}
}
