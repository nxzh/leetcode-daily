package regular

import "fmt"

func climbStairs2(n int) int {
	c := make(map[int]int)
	var climbS func(n int) int
	climbS = func(n int) int {
		if n == 0 || n == 1 {
			return 1
		}
		if _, presents := c[n-1]; !presents {
			c[n-1] = climbS(n - 1)
		}
		if _, presents := c[n-2]; !presents {
			c[n-2] = climbS(n - 2)
		}
		return c[n-1] + c[n-2]
	}
	return climbS(n)
}

func T70() {
	k := climbStairs2(2)
	fmt.Println(k)
}
