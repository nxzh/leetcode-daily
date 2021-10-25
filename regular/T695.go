package regular

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				queue := [][]int{{i, j}}
				area := 0
				for len(queue) > 0 {
					area++
					cx, cy := queue[0][0], queue[0][1]
					queue = queue[1:]
					cells := [][]int{{cx - 1, cy}, {cx + 1, cy}, {cx, cy - 1}, {cx, cy + 1}}
					for _, c := range cells {
						x, y := c[0], c[1]
						if x < 0 || y < 0 || x == m || y == n || grid[x][y] == 0 {
							continue
						}
						grid[x][y] = 0
						queue = append(queue, c)
					}
					if max < area {
						max = area
					}
				}
			}
		}
	}
	return max
}
func maxAreaOfIsland1(grid [][]int) int {
	// for every cell
	// if it is 1 then set its value to zero, and calculate all the cells nearby
	// set the max area

	max := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			area := calculateArea(grid, i, j)
			if max < area {
				max = area
			}
		}
	}
	return max
}

func calculateArea(grid [][]int, i int, j int) int {
	if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return 1 + calculateArea(grid, i-1, j) + calculateArea(grid, i+1, j) + calculateArea(grid, i, j-1) + calculateArea(grid, i, j+1)
}
