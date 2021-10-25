package regular

import "fmt"

func orangesRotting(grid [][]int) int {
	min := 0
	n1 := 0
	n2 := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				n1++
				isolated := true
				if i > 0 {
					isolated = isolated && grid[i-1][j] == 0
				}
				if i < len(grid)-1 {
					isolated = isolated && grid[i+1][j] == 0
				}
				if j > 0 {
					isolated = isolated && grid[i][j-1] == 0
				}
				if j < len(grid[0])-1 {
					isolated = isolated && grid[i][j+1] == 0
				}
				if isolated {
					return -1
				}
			}
			if grid[i][j] == 2 {
				n2++
			}
		}
	}
	if n1 == 0 {
		return 0
	}
	if n2 == 0 {
		return -1
	}
	for {
		toChange := [][]int{}
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[0]); j++ {
				if (grid[i][j]) == 2 {
					if i > 0 && grid[i-1][j] == 1 {
						toChange = append(toChange, []int{i - 1, j})
					}
					if i < len(grid)-1 && grid[i+1][j] == 1 {
						toChange = append(toChange, []int{i + 1, j})
					}
					if j > 0 && grid[i][j-1] == 1 {
						toChange = append(toChange, []int{i, j - 1})
					}
					if j < len(grid[0])-1 && grid[i][j+1] == 1 {
						toChange = append(toChange, []int{i, j + 1})
					}
				}
			}
		}
		if len(toChange) > 0 {
			min++
			for i := 0; i < len(toChange); i++ {
				if grid[toChange[i][0]][toChange[i][1]] == 1 {
					grid[toChange[i][0]][toChange[i][1]] = 2
					n1--
				}
			}
			toChange = [][]int{}
		} else {
			if n1 == 0 {
				return min
			}
			return -1
		}
	}
	return min
}

func T994() {
	res := orangesRotting([][]int{{1, 1, 1, 1}})
	fmt.Println(res)
}
