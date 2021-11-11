package regular

import "fmt"

func permute(nums []int) [][]int {
	result := [][]int{}
	permuteInternal(nums, 0, &result)
	return result
}

func permuteInternal(nums []int, from int, result *[][]int) {
	if from == len(nums) {
		copied := make([]int, len(nums))
		copy(copied, nums)
		*result = append(*result, copied)
	}
	for j := from; j < len(nums); j++ {
		swapNums(nums, from, j)
		permuteInternal(nums, from+1, result)
		swapNums(nums, from, j)
	}
}

func swapNums(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func permute1(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	result := [][]int{}
	queue := [][]int{{nums[0]}}
	for i := 1; i < len(nums); i++ {
		length := len(queue)
		for j := 0; j < length; j++ {
			cur := queue[j]
			for k := 0; k <= len(cur); k++ {
				inserted := insertToPos(k, nums[i], cur)

				if len(inserted) == len(nums) {
					result = append(result, inserted)
				} else {
					queue = append(queue, inserted)
				}
			}
		}
		queue = queue[length:]
	}
	return result
}

func insertToPos(pos int, data int, elements []int) []int {
	ret := make([]int, len(elements)+1)
	for i := 0; i < pos; i++ {
		ret[i] = elements[i]
	}
	ret[pos] = data
	for i := pos + 1; i < len(ret); i++ {
		ret[i] = elements[i-1]
	}
	return ret
}

func T46() {
	res := permute([]int{1, 2, 3})
	fmt.Println(res)
}
