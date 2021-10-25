package regular

import (
	"fmt"
	"regexp"
)

func maxConsecutiveAnswers(answerKey string, k int) int {
	// separate by F+, get TS array. separate by T+, get FS array.
	// compose a new array counts from TS, FS.
	//		each elem is (len(TS[0]), len(FS[0]), len(TS[1], len(FS[1],...) if answerKey[0] is T.
	//		or (len(FS[0]), len(TS[0]), len(FS[1], len(TS[1],...) if answerKey[1] is F.
	// iterate counts with two points....
	re := regexp.MustCompile("T+")
	split := re.Split(answerKey, -1)
	var FS []string
	for i := range split {
		if len(split[i]) > 0 {
			FS = append(FS, split[i])
		}
	}
	re = regexp.MustCompile("F+")
	split = re.Split(answerKey, -1)
	var TS []string
	for i := range split {
		if len(split[i]) > 0 {
			TS = append(TS, split[i])
		}
	}
	var first []string
	var second []string
	counts := make([]int, len(FS)+len(TS))
	if answerKey[0] == 'T' {
		first, second = TS, FS
	} else {
		first, second = FS, TS
	}
	i, j, k := 0, 0, 0
	for {
		if i < len(first) {
			counts[k] = len(first[i])
			i++
			k++
		}
		if j < len(second) {
			counts[k] = len(second[j])
			j++
			k++
		}
		if i == len(first) && j == len(second) {
			break
		}
	}
	return 0
}

func T5873() {
	res := maxConsecutiveAnswers("TTFF", 2)
	//res := maxConsecutiveAnswers("TFFT", 1)
	//res := maxConsecutiveAnswers("TTFTTFTT", 1)
	fmt.Println(res)
}
