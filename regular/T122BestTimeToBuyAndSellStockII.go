package regular

func maxProfit2(prices []int) int {
	max := 0
	for i := 1; i < len(prices); i++ {
		dist := prices[i] - prices[i-1]
		if dist > 0 {
			max += dist
		}
	}
	return max
}
