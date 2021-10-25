package regular

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	positive := true
	begin := 0
	if s[0] == '-' || s[0] == '+' {
		begin = 1
		if s[0] == '-' {
			positive = false
		}
	}
	for ; begin < len(s); begin++ {
		if s[begin] != '0' {
			break
		}
	}
	end := begin
	for ; end < len(s); end++ {
		if isNum(s[end]) {
			continue
		} else {
			break
		}
	}
	if end-begin > 10 {
		if positive {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}
	sum := int32(0)
	for i := begin; i < end; i++ {
		sum = sum*10 + int32(s[i]-'0')
		if sum >= math.MaxInt32/10 && i+1 < end {
			if sum > math.MaxInt32/10 {
				if positive {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}
			digit := s[i+1] - '0'
			if positive {
				if digit >= 7 {
					return math.MaxInt32
				} else {
					return int(sum*10 + int32(digit))
				}
			} else {
				if digit >= 8 {
					return math.MinInt32
				} else {
					return int(sum*10+int32(digit)) * -1
				}
			}
		}
	}
	if positive {
		return int(sum)
	}
	return int(sum * -1)
}

func isNum(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func T08() {
	//fmt.Println(math.MinInt32)
	res := myAtoi("2147483646")
	fmt.Println(res)
}
