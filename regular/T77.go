package regular

import "fmt"

func combine(n int, k int) [][]int {
	queue := [][]int{{}}
	results := [][]int{}
	for i := 1; i <= n; i++ {
		size := len(queue)
		for j := 0; j < size; j++ {
			copied := make([]int, len(queue[j]))
			copy(copied, queue[j])
			copied = append(copied, i)
			if len(copied) == k {
				results = append(results, copied)
			} else if len(copied) < k {
				queue = append(queue, copied)
			}
		}
	}
	return results
}

func T77() {
	res := combine(20, 10)
	fmt.Println(res)
}
