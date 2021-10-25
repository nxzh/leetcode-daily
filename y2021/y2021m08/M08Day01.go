package y2021m08

import (
	"fmt"
	"math"
)

var max = 0

func largestIsland(grid [][]int) int {
	cache := make([][]int, len(grid[0]))
	for i := range cache {
		cache[i] = make([]int, len(grid[0]))
	}
	island(0, 0, len(grid), false, grid, cache)
	return max
}

func island(i, j, n int, used bool, grid [][]int, cache [][]int) int {
	if i < 0 || i >= n || j < 0 || j >= n {
		fmt.Printf("a %d %d : %d \n", i, j, 0)
		return 0
	}
	if cache[i][j] == 1 {
		return 0
	}
	cache[i][j] = 1
	if grid[i][j] == 1 {
		down := island(i+1, j, n, used, grid, cache)
		right := island(i, j+1, n, used, grid, cache)
		left := island(i, j-1, n, used, grid, cache)
		top := island(i-1, j, n, used, grid, cache)
		fmt.Printf("b. %d %d : %d \n", i, j, 1+down+right+left+top)
		res := 1 + down + right + left + top
		max = int(math.Max(float64(max), float64(res)))
		return res
	} else {
		if used {
			fmt.Printf("%d %d : %d \n", i, j, 0)
			return 0
		} else {
			newCache := make([][]int, len(grid[0]))
			for i := range newCache {
				newCache[i] = make([]int, len(grid[0]))
			}
			resNext := island(i+1, j+1, n, false, grid, newCache)
			down := island(i+1, j, n, true, grid, cache)
			right := island(i, j+1, n, true, grid, cache)
			up := island(i-1, j, n, true, grid, cache)
			left := island(i, j-1, n, true, grid, cache)
			res := 1 + down + right + left + up
			fmt.Printf("%d %d : %d \n", i, j, res)
			max = int(math.Max(float64(resNext), float64(res)))
			return res
		}
	}
}
