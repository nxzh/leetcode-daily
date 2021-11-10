package regular

import "fmt"

func combine(n int, k int) [][]int {
	queue := [][]int{{}}
	result := [][]int{}
	for i := 1; i <= n; i++ {
		length := len(queue)
		for j := 0; j < length; j++ {
			copied := make([]int, len(queue[j]))
			copy(copied, queue[j])
			copied = append(copied, i)
			if len(copied) == k {
				result = append(result, copied)
			} else {
				queue = append(queue, copied)
			}
		}
	}
	return result
}

func combine1(n int, k int) [][]int {
	// set a list to put the result
	// set a list to put the path
	result := [][]int{}
	path := []int{}
	combineInternal(1, n, k, path, &result)
	return result
}

func combineInternal(begin int, n int, k int, path []int, result *[][]int) {
	// if the path.length == k, add to the result
	if len(path) == k {
		copied := []int{}
		copied = append(copied, path...)
		*result = append(*result, copied)
		return
	}
	// from begin to n, for every element e
	// add the e to the path
	// call combineInternal(begin+1, n, path, list) again from the next element
	// remove e from the path
	for i := begin; i <= n; i++ {
		if n-i+1 < k-len(path) {
			break
		}
		path = append(path, i)
		combineInternal(i+1, n, k, path, result)
		path = path[0 : len(path)-1]
	}
}

func T77() {
	res := combine(4, 2)
	fmt.Println(res)
}
