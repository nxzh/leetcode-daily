package regular

func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		stLen := len(stack)
		last := byte(0)
		if stLen != 0 {
			last = stack[stLen-1]
		}
		if ch == '}' && last == '{' || ch == ']' && last == '[' || ch == ')' && last == '(' {
			stack = stack[:stLen-1]
		} else if ch == '{' || ch == '[' || ch == '(' {
			stack = append(stack, ch)
		} else {
			return false
		}
	}
	return len(stack) == 0
}
