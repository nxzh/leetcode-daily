package y2021m08

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	cnt := len(people)
	saved := make([]bool, cnt)
	count := 0
	for i := cnt - 1; i >= 0; i-- {
		if !saved[i] {
			saved[i] = true
			qualified := findCompany(people, 0, i-1, limit-people[i])
			for i := qualified; i >= 0; i-- {
				if !saved[i] {
					saved[i] = true
					break
				}
			}
			count++
		}
	}
	return count
}

func findCompany(people []int, from int, to int, target int) int {
	if to < from {
		return -1
	}
	if people[to] <= target {
		return to
	}
	if people[from] > target {
		return -1
	}
	low := from
	high := to
	for low <= high {
		mid := (low + high) / 2
		if people[mid] < target {
			low = mid + 1
		} else if people[mid] > target {
			if people[mid-1] <= target {
				return mid - 1
			}
			high = mid - 1
		} else {
			if mid+1 <= high && people[mid+1] > target {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

func Call0826CN() {
	//res := numRescueBoats([]int{3, 5, 3, 4}, 5)
	res := numRescueBoats([]int{
		8, 3, 8, 3, 10, 2, 9, 1, 3, 6, 6, 4, 2, 3, 3, 8, 10, 6, 1, 8, 4, 4, 6, 3, 10, 2, 5, 3, 6, 6, 7, 6, 5, 7, 5, 8, 8, 3, 4, 7, 2, 7, 4, 6, 2, 7, 4, 5, 5, 5, 7, 4, 7, 1, 4, 8, 1, 7, 1, 5, 9, 1, 6, 1, 9, 7, 8, 7, 1, 1, 7, 10, 9, 7, 8, 3, 8, 3, 2, 5, 4, 2, 5, 9, 5, 5, 8, 6, 2, 10, 5, 8, 4, 9, 4, 3, 2, 10, 6, 1}, 10)
	fmt.Println(res)
}
