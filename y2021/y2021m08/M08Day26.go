package y2021m08

import (
	"fmt"
	"regexp"
)

func isValidSerialization(preorder string) bool {
	m1 := regexp.MustCompile(`[0-9]+,#,#`)
	lastLen := len(preorder)
	for {
		preorder = m1.ReplaceAllString(preorder, "#")
		if preorder == "#" {
			return true
		}
		if lastLen == len(preorder) {
			return false
		}
		lastLen = len(preorder)
	}
	return false
}

func Call0826() {
	res := isValidSerialization("9,9,9,19,#,9,#,#,#,9,#,69,#,#,#")
	fmt.Println(res)
}
