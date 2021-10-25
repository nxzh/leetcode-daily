package regular

import "fmt"

func generate(numRows int) [][]int {
	res := [][]int{{1}}
	if numRows == 1 {
		return res
	}
	for i := 1; i < numRows; i++ {
		last := res[i-1]
		newRow := []int{1}
		for i := 0; i < len(last)-1; i++ {
			newRow = append(newRow, last[i]+last[i+1])
		}
		newRow = append(newRow, 1)
		res = append(res, newRow)
	}
	return res
}

func T118() {
	res := generate(5)
	fmt.Println(res)
}
