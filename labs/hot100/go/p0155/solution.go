package p0155

// 每个元素同时记下「压入它时的栈内最小值」，弹出后最小值自动回到上一状态。
type MinStack struct {
	stack []item
}

type item struct{ val, min int }

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	m := val
	if len(s.stack) > 0 {
		m = min(m, s.stack[len(s.stack)-1].min)
	}
	s.stack = append(s.stack, item{val, m})
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1].val
}

func (s *MinStack) GetMin() int {
	return s.stack[len(s.stack)-1].min
}
