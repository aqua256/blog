package p0160

import (
	"math/rand/v2"
	"testing"
)

// joined 构造 A = a + common, B = b + common，返回两头和交点。
func joined(a, b, common []int) (*ListNode, *ListNode, *ListNode) {
	c := fromSlice(common)
	attach := func(vals []int) *ListNode {
		h := fromSlice(vals)
		if h == nil {
			return c
		}
		p := h
		for p.Next != nil {
			p = p.Next
		}
		p.Next = c
		return h
	}
	return attach(a), attach(b), c
}

func TestExamples(t *testing.T) {
	cases := []struct{ a, b, c []int }{
		{[]int{4, 1}, []int{5, 6, 1}, []int{8, 4, 5}},
		{[]int{1, 9, 1}, []int{3}, []int{2, 4}},
		{[]int{2, 6, 4}, []int{1, 5}, nil},
	}
	for _, tc := range cases {
		ha, hb, want := joined(tc.a, tc.b, tc.c)
		if got := getIntersectionNode(ha, hb); got != want {
			t.Errorf("case %v wrong node", tc)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 160))
	gen := func(n int) []int {
		s := make([]int, n)
		for i := range s {
			s[i] = r.IntN(3) // 值大量重复，确保比的是节点而不是值
		}
		return s
	}
	for range 1000 {
		a, b, c := gen(r.IntN(8)), gen(r.IntN(8)), gen(r.IntN(6))
		if len(a)+len(c) == 0 || len(b)+len(c) == 0 {
			continue // 题目保证两表非空
		}
		ha, hb, want := joined(a, b, c)
		if got := getIntersectionNode(ha, hb); got != want {
			t.Fatalf("a=%v b=%v c=%v wrong node", a, b, c)
		}
	}
}
