package p0199

import (
	"math/rand/v2"
	"slices"
	"testing"
)

// 暴力：BFS 每层取最后一个
func brute(root *TreeNode) []int {
	res := []int{}
	level := []*TreeNode{}
	if root != nil {
		level = append(level, root)
	}
	for len(level) > 0 {
		res = append(res, level[len(level)-1].Val)
		var next []*TreeNode
		for _, n := range level {
			for _, c := range []*TreeNode{n.Left, n.Right} {
				if c != nil {
					next = append(next, c)
				}
			}
		}
		level = next
	}
	return res
}

func TestExamples(t *testing.T) {
	cases := []struct{ in, want []int }{
		{[]int{1, 2, 3, null, 5, null, 4}, []int{1, 3, 4}},
		{[]int{1, 2, 3, 4, null, null, null, 5}, []int{1, 3, 4, 5}},
		{[]int{1, null, 3}, []int{1, 3}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		if got := rightSideView(build(c.in)); !slices.Equal(got, c.want) {
			t.Errorf("%v: got %v want %v", c.in, got, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(199, 1))
	for range 500 {
		root := randTree(r, r.IntN(40), 0, 99)
		if got, want := rightSideView(root), brute(root); !slices.Equal(got, want) {
			t.Fatalf("%v: got %v want %v", serialize(root), got, want)
		}
	}
}
