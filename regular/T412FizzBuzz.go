package regular

import "strconv"

var m3 = make(map[int]bool, 3334)
var m5 = make(map[int]bool, 2000)

func init() {
	for i := 3; i <= 10000; i += 3 {
		m3[i] = true
	}
	for i := 5; i <= 10000; i += 5 {
		m5[i] = true
	}
}

func fizzBuzz(n int) []string {

	var res []string
	for i := 1; i <= n; i++ {
		if3 := m3[i]
		if5 := m5[i]
		if if3 && if5 {
			res = append(res, "FizzBuzz")
		} else if if3 {
			res = append(res, "Fizz")
		} else if if5 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}
