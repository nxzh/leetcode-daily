package regular

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	result := 0
	count := 0
	for x != 0 {
		count++
		remain := x % 10
		if count == 10 {
			if math.Abs(float64(result)) > 214748364 || remain < -8 || remain > 7 {
				return 0
			}
		}
		fmt.Println(remain)
		result = result*10 + remain
		x /= 10
	}
	return result
}

func main2() {
	fmt.Println(reverse(-321))
}
