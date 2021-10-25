package y2021m08

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		cache := make([]byte, 10)
		for j := 0; j < 9; j++ {
			k := board[i][j]
			if board[i][j] == '.' {
				continue
			}
			k -= '0'
			cache[k] += 1
			if cache[k] > 1 {
				return false
			}
		}
	}
	for i := 0; i < 9; i++ {
		cache := make([]byte, 10)
		for j := 0; j < 9; j++ {
			k := board[j][i]
			if k == '.' {
				continue
			}
			k -= '0'
			cache[k] += 1
			if cache[k] > 1 {
				return false
			}
		}
	}
	ki, kj := 1, 1
	for i := (ki - 1) * 3; ki <= 3; {
		cache := make([]byte, 10)
		for j := (kj - 1) * 3; j < kj*3; {
			k := board[i][j]
			if k != '.' {
				k -= '0'
				cache[k] += 1
				if cache[k] > 1 {
					return false
				}
			}
			if (j+1)%3 == 0 {
				if (i+1)%3 == 0 {
					break
				}
				i, j = i+1, (kj-1)*3
			} else {
				j++
			}
		}
		if (i+1)%3 == 0 {
			kj++
			i = (ki - 1) * 3
		}
		if kj > 3 {
			kj, ki = 1, ki+1
			i = (ki - 1) * 3
		}
	}
	return true
}
