package p0020

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	pair := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		open, isClose := pair[c]
		if !isClose {
			stack = append(stack, c)
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != open {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
