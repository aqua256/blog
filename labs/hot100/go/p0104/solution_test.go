package p0104

import (
	"math/rand/v2"
	"testing"
)

// 暴力：BFS 数层数
func bruteDepth(root *TreeNode) int {
	depth := 0
	for level := []*TreeNode{root}; ; depth++ {
		var next []*TreeNode
		for _, n := range level {
			if n != nil {
				next = append(next, n.Left, n.Right)
			}
		}
		if len(next) == 0 {
			return depth
		}
		level = next
	}
}

func TestExamples(t *testing.T) {
	if got := maxDepth(build([]int{3, 9, 20, null, null, 15, 7})); got != 3 {
		t.Errorf("got %d want 3", got)
	}
	if got := maxDepth(build([]int{1, null, 2})); got != 2 {
		t.Errorf("got %d want 2", got)
	}
	if got := maxDepth(nil); got != 0 {
		t.Errorf("got %d want 0", got)
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(104, 1))
	for range 500 {
		root := randTree(r, r.IntN(30), 0, 9)
		if got, want := maxDepth(root), bruteDepth(root); got != want {
			t.Fatalf("%v: got %d want %d", serialize(root), got, want)
		}
	}
}
