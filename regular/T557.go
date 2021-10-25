package regular

import "fmt"

func reverseWords(s string) string {
	l, r := 0, 0
	result := []byte{}
	for {
		for r != len(s) && s[r] != ' ' {
			r++
		}
		res := reverseSec(s, l, r)
		result = append(result, res...)
		l = r
		if l >= len(s) {
			return string(result)
		}
		for s[l] == ' ' {
			result = append(result, s[l])
			l++
		}
		r = l
	}
}

func reverseSec(str string, from int, to int) []byte {
	s := make([]byte, to-from)
	for i := 0; i < len(s); i++ {
		s[i] = str[to-i-1]
	}
	return s
}

func T557() {
	res := reverseWords("hello, world!")
	fmt.Println(res)
}
