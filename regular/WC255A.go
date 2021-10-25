package regular

func findGCD(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minNum, maxNum := nums[0], 0
	for _, num := range nums {
		if minNum > num {
			minNum = num
		}
		if maxNum < num {
			maxNum = num
		}
	}
	return gcd(maxNum, minNum)
}

func gcd(a, b int) int {
	if b <= 0 {
		return a
	}
	return gcd(b, a%b)
}
