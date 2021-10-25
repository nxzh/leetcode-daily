package regular

func titleToNumber(columnTitle string) int {
	sum := 0
	for i := 0; i < len(columnTitle); i++ {
		sum = sum*26 + int(letterToI(columnTitle[i]))
	}

	return sum
}

func letterToI(ch uint8) uint8 {
	return ch - 'A' + 1
}
