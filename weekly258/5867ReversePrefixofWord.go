package weekly258

import "fmt"

func reversePrefix(word string, ch byte) string {
	bts := []byte(word)
	pos := -1
	for i := 0; i < len(bts); i++ {
		if bts[i] == ch {
			pos = i + 1
			break
		}
	}
	if pos == -1 {
		return string(bts)
	}
	for i := 0; i < pos/2; i++ {
		temp := bts[i]
		bts[i] = bts[pos-i-1]
		bts[pos-i-1] = temp
	}
	return string(bts)
}

func CallW258() {
	res := reversePrefix("abcdefd", 'd')
	fmt.Println(res)
}
