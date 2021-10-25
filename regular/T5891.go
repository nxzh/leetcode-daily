package regular

import "fmt"

func missingRolls(rolls []int, mean int, n int) []int {
	sumM := 0
	lenM := len(rolls)
	for i := 0; i < lenM; i++ {
		sumM += rolls[i]
	}
	total := mean * (lenM + n)
	sumN := total - sumM
	if sumN < n || sumN > 6*n {
		return []int{}
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 6
	}
	i := 0
	for cur := 6 * n; cur > sumN; cur-- {
		res[i]--
		if res[i] == 1 {
			i += 1
		}
	}
	return res
}

func T5891() {
	res := missingRolls([]int{1, 2, 3, 4}, 6, 4)
	fmt.Println(res)
}
