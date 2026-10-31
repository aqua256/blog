package p0141

import (
	"math/rand/v2"
	"testing"
)

// withCycle 构造链表，尾部指向下标 pos（-1 表示无环）。
func withCycle(vals []int, pos int) *ListNode {
	head := fromSlice(vals)
	if pos < 0 {
		return head
	}
	var entry, tail *ListNode
	for i, p := 0, head; p != nil; i, p = i+1, p.Next {
		if i == pos {
			entry = p
		}
		tail = p
	}
	tail.Next = entry
	return head
}

func TestExamples(t *testing.T) {
	cases := []struct {
		vals []int
		pos  int
		want bool
	}{
		{[]int{3, 2, 0, -4}, 1, true},
		{[]int{1, 2}, 0, true},
		{[]int{1}, -1, false},
		{[]int{}, -1, false},
	}
	for _, c := range cases {
		if got := hasCycle(withCycle(c.vals, c.pos)); got != c.want {
			t.Errorf("hasCycle(%v,pos=%d) = %v", c.vals, c.pos, got)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 141))
	for range 500 {
		n := r.IntN(15)
		vals := make([]int, n)
		pos := -1
		if n > 0 && r.IntN(2) == 0 {
			pos = r.IntN(n)
		}
		if got := hasCycle(withCycle(vals, pos)); got != (pos >= 0) {
			t.Fatalf("n=%d pos=%d got %v", n, pos, got)
		}
	}
}
