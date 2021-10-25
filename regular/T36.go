package regular

import "fmt"

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		buffer := make([]byte, 10)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			index := board[i][j] - '0'
			buffer[index] += 1
			if buffer[index] > 1 {
				return false
			}
		}
	}
	for i := 0; i < 9; i++ {
		buffer := make([]byte, 10)
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			index := board[j][i] - '0'
			buffer[index] += 1
			if buffer[index] > 1 {
				return false
			}
		}
	}
	c, r := 0, 0
	for r < 9 {
		buffer := make([]byte, 10)
		for i := r; i < r+3; i++ {
			for j := c; j < c+3; j++ {
				if board[i][j] == '.' {
					continue
				}
				index := board[i][j] - '0'
				buffer[index] += 1
				if buffer[index] > 1 {
					return false
				}
			}
		}
		c += 3
		if c == 9 {
			c = 0
			r += 3
		}
	}
	return true
}

func T36() {
	res := isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	})
	fmt.Println(res)
}
