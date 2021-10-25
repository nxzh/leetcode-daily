package y2021m08

import (
	"fmt"
	"math"
)

func minPatches(nums []int, n int) int {
	filledMap := make(map[int]bool)
	filled := []int{0}
	sumOfSet := func(aSet []int) (sum int) {
		for i := 0; i < len(aSet); i++ {
			sum += aSet[i]
		}
		return sum
	}
	curSet := [][]int{{}}
	for i := 0; i < len(nums); i++ {
		curSetLen := len(curSet)
		for j := 0; j < curSetLen; j++ {
			tryAdding := append(curSet[j], nums[i])
			sumOfTryAdding := sumOfSet(tryAdding)
			if sumOfTryAdding <= n {
				if !filledMap[sumOfTryAdding] {
					filledMap[sumOfTryAdding] = true
					filled = append(filled, sumOfTryAdding)
				}
				curSet = append(curSet, tryAdding)
			}
		}
	}

	minKey := func(aMap map[int]bool) int {
		minK := math.MaxInt32
		for key := range aMap {
			if aMap[key] {
				if minK > key {
					minK = key
				}
			}

		}
		return minK
	}
	count := 0
	for len(filled) <= n {
		//fmt.Println(filled)
		minK := minKey(filledMap)
		count++
		filledLen := len(filled)
		for i := 0; i < filledLen; i++ {
			tryAdding := filled[i] + minK
			if tryAdding <= n {
				if !filledMap[tryAdding] {
					filledMap[tryAdding] = true
					filled = append(filled, tryAdding)
				}
			}
		}
	}
	return count
}

func Call0829() {
	res := minPatches([]int{1, 2, 31, 33}, 2147483647)
	fmt.Println(res)
}
