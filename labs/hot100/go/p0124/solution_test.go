package p0124

import (
	"math"
	"math/rand/v2"
	"testing"
)

// 暴力 O(n^2)：建无向图，从每个点出发 DFS，所有简单路径的和取最大
func brute(root *TreeNode) int {
	adj := map[*TreeNode][]*TreeNode{}
	for _, n := range nodes(root) {
		for _, c := range []*TreeNode{n.Left, n.Right} {
			if c != nil {
				adj[n] = append(adj[n], c)
				adj[c] = append(adj[c], n)
			}
		}
	}
	best := math.MinInt
	var dfs func(u, from *TreeNode, s int)
	dfs = func(u, from *TreeNode, s int) {
		s += u.Val
		best = max(best, s)
		for _, v := range adj[u] {
			if v != from {
				dfs(v, u, s)
			}
		}
	}
	for _, n := range nodes(root) {
		dfs(n, nil, 0)
	}
	return best
}

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{-10, 9, 20, null, null, 15, 7}, 42},
		{[]int{-3}, -3},
		{[]int{2, -1}, 2},
	}
	for _, c := range cases {
		if got := maxPathSum(build(c.in)); got != c.want {
			t.Errorf("%v: got %d want %d", c.in, got, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(124, 1))
	for range 500 {
		root := randTree(r, 1+r.IntN(30), -10, 10)
		if got, want := maxPathSum(root), brute(root); got != want {
			t.Fatalf("%v: got %d want %d", serialize(root), got, want)
		}
	}
}
