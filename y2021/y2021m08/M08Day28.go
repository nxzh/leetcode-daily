package y2021m08

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	dp := make(map[int]map[string]int, len(startTime))
	for i := 0; i < len(startTime); i++ {
		dp[i] = make(map[string]int)
	}
	minStart := findMinin(startTime)
	maxEnd := findMax(endTime)
	return jobSchedulingInternal(dp, startTime, endTime, profit, 0, &[]int{minStart, maxEnd})
}
func findMinin(elems []int) int {
	min := math.MaxInt32
	for _, e := range elems {
		if min > e {
			min = e
		}
	}
	return min
}
func findMax(elems []int) int {
	max := 0
	for _, e := range elems {
		if max < e {
			max = e
		}
	}
	return max
}
func jobSchedulingInternal(dp map[int]map[string]int, startTime []int, endTime []int, profit []int, currentIndex int, remainTimeSecs *[]int) int {
	if currentIndex >= len(startTime) {
		return 0
	}
	toKey := func(nums []int) string {
		keyStrs := []string{}
		for i := 0; i < len(nums); i++ {
			keyStrs = append(keyStrs, strconv.Itoa(nums[i]))
		}
		return strings.Join(keyStrs, ",")
	}
	key := toKey(*remainTimeSecs)
	if _, prs := dp[currentIndex][key]; !prs {

		remainTimeSecsCopy1 := make([]int, len(*remainTimeSecs))
		var remainTimeSecsCopy2 []int
		copy(remainTimeSecsCopy1, *remainTimeSecs)
		profitSkipCurrentIndex := jobSchedulingInternal(dp, startTime, endTime, profit, currentIndex+1, &remainTimeSecsCopy1)
		profitAddCurrentIndex := 0
		for i := 0; i < len(*remainTimeSecs); i += 2 {
			if (*remainTimeSecs)[i] <= startTime[currentIndex] && (*remainTimeSecs)[i+1] >= endTime[currentIndex] {
				remainTimeSecsCopy2 = make([]int, len(*remainTimeSecs))
				copy(remainTimeSecsCopy2, *(remainTimeSecs))
				remainTimeSecsCopy2 = append(remainTimeSecsCopy2[:i], remainTimeSecsCopy2[i+2:]...)
				sec1 := []int{(*remainTimeSecs)[i], startTime[currentIndex]}
				sec2 := []int{endTime[currentIndex], (*remainTimeSecs)[i+1]}
				k := i
				if sec1[0] != sec1[1] {
					remainTimeSecsCopy2 = append(remainTimeSecsCopy2[:k], sec1...)
					k += 2
				}
				if sec2[0] != sec2[1] {
					remainTimeSecsCopy2 = append(remainTimeSecsCopy2[:k], sec2...)
					k += 2
				}
				remainTimeSecsCopy2 = append(remainTimeSecsCopy2[:k], (*remainTimeSecs)[i+2:]...)
				//fmt.Println("-> Add index: ", currentIndex, "[", startTime[currentIndex], endTime[currentIndex], "]", "to -> [", toKey(*remainTimeSecs), "]", "->", profitAddCurrentIndex)

				profitAddCurrentIndex = profit[currentIndex] + jobSchedulingInternal(dp, startTime, endTime, profit, currentIndex+1, &remainTimeSecsCopy2)
				break
			}
		}
		if profitSkipCurrentIndex >= profitAddCurrentIndex {
			//fmt.Println("-> Result Of skip index: ", currentIndex, "[", startTime[currentIndex], endTime[currentIndex], "]", "to -> [", toKey(*remainTimeSecs), "]", "->", profitSkipCurrentIndex)
			*remainTimeSecs = remainTimeSecsCopy1
			//dp[currentIndex][key] = profitSkipCurrentIndex
			return profitSkipCurrentIndex
		} else {
			//fmt.Println("-> Result of add index: ", currentIndex, "[", startTime[currentIndex], endTime[currentIndex], "]", "to -> [", toKey(*remainTimeSecs), "]", "->", profitAddCurrentIndex)
			*remainTimeSecs = remainTimeSecsCopy2
			//dp[currentIndex][key] = profitAddCurrentIndex
			return profitAddCurrentIndex
		}
	} else {
		return dp[currentIndex][key]
	}
	return dp[currentIndex][key]
}

func Call0828() {
	//res := jobScheduling([]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4})
	//res := jobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70})
	res := jobScheduling([]int{43, 13, 36, 31, 40, 5, 47, 13, 28, 16, 2, 11}, []int{44, 22, 41, 41, 47, 13, 48, 35, 48, 26, 21, 39}, []int{8, 20, 3, 19, 16, 8, 11, 13, 2, 15, 1, 1})
	fmt.Println(res)
}
