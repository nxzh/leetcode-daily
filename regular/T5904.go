package regular

import "fmt"

func countMaxOrSubsets(nums []int) int {
	max := 0
	maxCount := 0
	queue := [][]int{{}}
	for _, num := range nums {
		size := len(queue)
		for i := 0; i < size; i++ {
			// TODO: copy queue[i], don't use the origin element.
			elem := append(queue[i], num)
			bitOr := 0
			for j := 0; j < len(elem); j++ {
				bitOr = bitOr | elem[j]
			}
			if max == bitOr {
				maxCount++
			} else if max < bitOr {
				max = bitOr
				maxCount = 1
			}
			queue = append(queue, elem)
		}
	}
	return maxCount
}
func T5904() {

	max := countMaxOrSubsets([]int{4, 4, 5, 5, 6, 6, 7, 8})
	fmt.Println(max)
}
