package y2021m08

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {

	part1 := 0

	for {
		part2 := c - part1*part1
		if part2 < part1 {
			return false
		}

		root := math.Sqrt(float64(part2))
		if root-float64(int(root)) == 0 {
			return true
		}

		part1++
	}
}

func Call0825() {
	res := judgeSquareSum(6)
	fmt.Println(res)
}
