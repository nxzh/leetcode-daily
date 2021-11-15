package regular

import (
	"fmt"
	"math"
)

func checkAlmostEquivalent(word1 string, word2 string) bool {
	var ca [26]int
	for i := 0; i < len(word1); i++ {
		ca[word1[i]-'a'] += 1
	}
	var cb [26]int
	for i := 0; i < len(word2); i++ {
		cb[word2[i]-'a'] += 1
	}
	for i := 0; i < 26; i++ {
		if math.Abs(float64(ca[i]-cb[i])) > 3 {
			return false
		}
	}
	return true
}

func T5910() {
	fmt.Println(checkAlmostEquivalent("cccddabba", "babababab"))
}
