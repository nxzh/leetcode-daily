package regular

import (
	"fmt"
	"strings"
)

func letterCasePermutation(s string) []string {
	queue := []string{""}
	var f func(s string, next int)
	f = func(s string, next int) {
		if next == len(s) {
			return
		}
		size := len(queue)
		if s[next] >= 'a' && s[next] <= 'z' || s[next] >= 'A' && s[next] <= 'Z' {
			for i := 0; i < size; i++ {
				e1 := queue[i] + string(s[next])
				queue = append(queue, e1)
				e2 := queue[i] + string(s[next]+'A'-'a')
				queue = append(queue, e2)
			}
		} else {
			for i := 0; i < size; i++ {
				e := queue[i] + string(s[next])
				queue = append(queue, e)
			}
		}
		queue = queue[size:]
		f(s, next+1)
	}
	f(strings.ToLower(s), 0)
	return queue
}

func T784() {
	res := letterCasePermutation("C")
	fmt.Println(res)
}
