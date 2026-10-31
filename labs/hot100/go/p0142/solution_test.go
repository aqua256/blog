package p0142

import (
	"math/rand/v2"
	"testing"
)

// withCycle 构造链表，尾部指向下标 pos；返回头和入环节点（无环为 nil）。
func withCycle(vals []int, pos int) (*ListNode, *ListNode) {
	head := fromSlice(vals)
	if pos < 0 {
		return head, nil
	}
	var entry, tail *ListNode
	for i, p := 0, head; p != nil; i, p = i+1, p.Next {
		if i == pos {
			entry = p
		}
		tail = p
	}
	tail.Next = entry
	return head, entry
}

func TestExamples(t *testing.T) {
	cases := []struct {
		vals []int
		pos  int
	}{
		{[]int{3, 2, 0, -4}, 1},
		{[]int{1, 2}, 0},
		{[]int{1}, -1},
	}
	for _, c := range cases {
		head, want := withCycle(c.vals, c.pos)
		if got := detectCycle(head); got != want {
			t.Errorf("detectCycle(%v,pos=%d) wrong node", c.vals, c.pos)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 142))
	for range 1000 {
		n := r.IntN(20)
		vals := make([]int, n)
		pos := -1
		if n > 0 && r.IntN(4) != 0 {
			pos = r.IntN(n)
		}
		head, want := withCycle(vals, pos)
		if got := detectCycle(head); got != want {
			t.Fatalf("n=%d pos=%d wrong node", n, pos)
		}
	}
}
