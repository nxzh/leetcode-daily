package weekly258

import "fmt"

func interchangeableRectangles(rectangles [][]int) int64 {
	ratios := make(map[float64]int64)
	for i := 0; i < len(rectangles); i++ {
		ratio := float64(rectangles[i][0]) / float64(rectangles[i][1])
		ratios[ratio]++
	}
	output := int64(0)
	for _, v := range ratios {
		if v > 1 {
			output += v * (v - 1) / 2
		}
	}
	return output
}

func Call5868() {
	res := interchangeableRectangles([][]int{{4, 8}, {3, 6}, {10, 20}, {15, 30}})
	fmt.Println(res)
}
