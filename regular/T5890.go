package regular

import "fmt"

func minimumMoves(s string) int {
	nextChar := func(from int, ch byte) int {
		for i := from; i < len(s); i++ {
			if s[i] == ch {
				return i
			}
		}
		return len(s)
	}

	count := 0

	l := nextChar(0, 'X')
	for {
		if l >= len(s) {
			break
		}
		r := nextChar(l, 'O')
		n := (r - l) / 3
		if (r-l)%3 != 0 {
			n += 1
		}
		count += n
		l = n*3 + l
		l = nextChar(l, 'X')
	}
	return count
}

func T5890() {
	res := minimumMoves("OOOO")
	fmt.Println(res)
}
