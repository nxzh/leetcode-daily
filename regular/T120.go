package regular

import (
	"fmt"
	"math"
)

func minimumTotal(triangle [][]int) int {
	length := len(triangle)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, i+1)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < i+1; j++ {
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == i {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i-1][j-1]))) + triangle[i][j]
			}
		}
	}
	min := math.MaxInt32
	for i := 0; i < length; i++ {
		if min > dp[length-1][i] {
			min = dp[length-1][i]
		}
	}
	return min
}

func T120() {
	n := minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})
	fmt.Println(n)
}
