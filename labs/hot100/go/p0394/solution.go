package p0394

import "strings"

// 遇到 '[' 把「当前已拼好的串」和「重复次数」压栈，开始新的一层；
// 遇到 ']' 弹出外层，外层串 + 当前串×次数 成为新的当前串。
func decodeString(s string) string {
	type frame struct {
		prefix string
		k      int
	}
	stack := []frame{}
	cur := strings.Builder{}
	k := 0
	for _, c := range s {
		switch {
		case c >= '0' && c <= '9':
			k = k*10 + int(c-'0') // 次数可能是多位数
		case c == '[':
			stack = append(stack, frame{cur.String(), k})
			cur.Reset()
			k = 0
		case c == ']':
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			inner := strings.Repeat(cur.String(), top.k)
			cur.Reset()
			cur.WriteString(top.prefix + inner)
		default:
			cur.WriteRune(c)
		}
	}
	return cur.String()
}
