package regular

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	// BFS
	// solution733(image, oldColor, sr, sc, newColor)
	queue := [][]int{{sr, sc}}
	image[sr][sc] = newColor
	rows := len(image)
	cols := len(image[0])
	for len(queue) > 0 {
		// get the first element
		// change the color
		// and check the four direction cells
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		cells := [][]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
		for _, cell := range cells {
			if cell[0] >= 0 && cell[1] >= 0 && cell[0] < rows && cell[1] < cols && image[cell[0]][cell[1]] == oldColor {
				queue = append(queue, cell)
				image[cell[0]][cell[1]] = newColor
			}
		}
	}
	return image
}

func solution733(image [][]int, oldColor int, x int, y int, newColor int) {
	if x < 0 || y < 0 || x == len(image) || y == len(image[0]) || image[x][y] != oldColor {
		return
	}
	image[x][y] = newColor
	solution733(image, oldColor, x+1, y, newColor)
	solution733(image, oldColor, x-1, y, newColor)
	solution733(image, oldColor, x, y+1, newColor)
	solution733(image, oldColor, x, y-1, newColor)
}
