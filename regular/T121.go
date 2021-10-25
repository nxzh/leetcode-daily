package regular

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	// 1. from 1 to len-1
	// 2. dist = prices[i]-prices[i-1]
	// profit += dist
	// 3. if profit <  0 profit =0
	// 4. if profit > 0 profit += dist set max
	max, profit := 0, 0
	for i := 1; i < len(prices); i++ {
		dist := prices[i] - prices[i-1]
		profit += dist
		if profit < 0 {
			profit = 0
		} else {
			max = int(math.Max(float64(max), float64(profit)))
		}
	}
	return max
}

func T121() {
	res := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(res)
}
