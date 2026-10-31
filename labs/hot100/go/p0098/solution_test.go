package p0098

import (
	"math/rand/v2"
	"testing"
)

// 暴力：BST <=> 中序严格递增
func brute(root *TreeNode) bool {
	var vals []int
	var dfs func(*TreeNode)
	dfs = func(n *TreeNode) {
		if n != nil {
			dfs(n.Left)
			vals = append(vals, n.Val)
			dfs(n.Right)
		}
	}
	dfs(root)
	for i := 1; i < len(vals); i++ {
		if vals[i] <= vals[i-1] {
			return false
		}
	}
	return true
}

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want bool
	}{
		{[]int{2, 1, 3}, true},
		{[]int{5, 1, 4, null, null, 3, 6}, false},
		{[]int{5, 4, 6, null, null, 3, 7}, false}, // 3 在右子树却小于根
		{[]int{2, 2, 2}, false},
		{[]int{-2147483648, null, 2147483647}, true},
	}
	for _, c := range cases {
		if got := isValidBST(build(c.in)); got != c.want {
			t.Errorf("%v: got %v want %v", c.in, got, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(98, 1))
	trues := 0
	for i := range 1000 {
		var root *TreeNode
		switch i % 3 {
		case 0:
			root = randTree(r, r.IntN(8), 0, 9)
		case 1:
			root = randBST(r, r.IntN(20))
		default:
			// BST 上随机改一个值，制造「只违反祖先约束」的情况
			root = randBST(r, 1+r.IntN(20))
			ns := nodes(root)
			ns[r.IntN(len(ns))].Val += r.IntN(11) - 5
		}
		want := brute(root)
		if want {
			trues++
		}
		if got := isValidBST(root); got != want {
			t.Fatalf("%v: got %v want %v", serialize(root), got, want)
		}
	}
	if trues < 200 || trues > 800 {
		t.Fatalf("unbalanced cases: %d trues", trues)
	}
}
