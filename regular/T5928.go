package regular

import (
	"fmt"
	"strings"
)

func decodeCiphertext(encodedText string, rows int) string {
	if rows == 1 {
		return encodedText
	}
	cols := len(encodedText) / rows
	matrix := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]byte, cols)
	}
	index := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = encodedText[index]
			index++
		}
	}
	result := []byte{}
	for i := 0; i <= cols; i++ {
		for j := 0; j < rows && j+i < cols; j++ {
			result = append(result, matrix[j][j+i])
		}
	}
	return strings.TrimRight(string(result), " ")
}

func T5928() {
	res := decodeCiphertext("osljjaooouqphokrnf     lsdoioccbdhbsqkm b     qmqoitpqnpqwnebsou     llvfotmazegriuigrs     vjpbgaqifwo  kaqto     dsupahycdgbyoubsu",
		6)
	fmt.Println(res)
}
