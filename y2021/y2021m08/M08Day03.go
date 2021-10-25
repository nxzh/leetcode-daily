package y2021m08

import (
	"reflect"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{{}}
	for i := 0; i < len(nums); i++ {
		//res = append(res, []int{nums[i]})
		addToResult(&[]int{nums[i]}, &res)
		var elems [][]int
		for _, num := range nums {
			elems = append(elems, []int{num})
		}
		k := i
		for k < len(nums) {
			elem := elems[k]
			for j := k + 1; j < len(nums); j++ {
				elems[j] = append(elem, nums[j])
				addToResult(&elems[j], &res)
				//res = append(res, elems[j])
			}
			k++
		}
	}

	return res
}

func addToResult(elems *[]int, result *[][]int) {
	length := len(*elems)
	for i := 0; i < len(*result); i++ {
		if len((*result)[i]) == length {
			sort.Ints(*elems)
			sort.Ints((*result)[i])
			if reflect.DeepEqual(*elems, (*result)[i]) {
				return
			}
		}
	}
	*result = append(*result, *elems)
}
