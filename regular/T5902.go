package regular

import (
	"fmt"
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	var tokens = strings.Split(s, " ")
	next := 0
	for _, token := range tokens {
		allDigit := true
		for _, ch := range token {
			if ch < '0' || ch > '9' {
				allDigit = false
				break
			}
		}
		if allDigit {
			num, _ := strconv.Atoi(token)
			if next >= num {
				return false
			}
			next = num
		}
	}
	return true
}

func T5902() {
	result := areNumbersAscending("4 5 11 26")
	fmt.Println(result)
}
