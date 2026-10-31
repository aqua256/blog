package p0437

import (
	"math/rand/v2"
	"testing"
)

// 暴力 O(n^2)：以每个节点为起点向下枚举
func brute(root *TreeNode, target int) int {
	var down func(n *TreeNode, s int) int
	down = func(n *TreeNode, s int) int {
		if n == nil {
			return 0
		}
		s += n.Val
		c := 0
		if s == target {
			c = 1
		}
		return c + down(n.Left, s) + down(n.Right, s)
	}
	total := 0
	for _, n := range nodes(root) {
		total += down(n, 0)
	}
	return total
}

func TestExamples(t *testing.T) {
	if got := pathSum(build([]int{10, 5, -3, 3, 2, null, 11, 3, -2, null, 1}), 8); got != 3 {
		t.Errorf("got %d want 3", got)
	}
	if got := pathSum(build([]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1}), 22); got != 3 {
		t.Errorf("got %d want 3", got)
	}
	if got := pathSum(build([]int{1000000000, 1000000000, null, 294967296, null, 1000000000, null, 1000000000, null, 1000000000}), 0); got != 0 {
		t.Errorf("got %d want 0", got)
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(437, 1))
	for range 500 {
		root := randTree(r, r.IntN(30), -3, 3)
		target := r.IntN(9) - 4
		if got, want := pathSum(root, target), brute(root, target); got != want {
			t.Fatalf("%v target=%d: got %d want %d", serialize(root), target, got, want)
		}
	}
}
