package regular

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {
	count := 0
	for {
		for i := 0; i < len(tickets); i++ {
			if tickets[i] != 0 {
				tickets[i] -= 1
				count++
			}
			if tickets[k] == 0 {
				return count
			}
		}
	}
}

func T5926() {
	//t := timeRequiredToBuy([]int{2, 3, 2}, 2)
	t := timeRequiredToBuy([]int{84, 49, 5, 24, 70, 77, 87, 8}, 3)
	fmt.Println(t)
}
