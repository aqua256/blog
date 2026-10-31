package p0543

import (
	"math/rand/v2"
	"testing"
)

// 暴力 O(n^2)：建无向图，从每个点 BFS，取最远距离（边数）
func brute(root *TreeNode) int {
	adj := map[*TreeNode][]*TreeNode{}
	for _, n := range nodes(root) {
		adj[n] = adj[n]
		for _, c := range []*TreeNode{n.Left, n.Right} {
			if c != nil {
				adj[n] = append(adj[n], c)
				adj[c] = append(adj[c], n)
			}
		}
	}
	best := 0
	for s := range adj {
		dist := map[*TreeNode]int{s: 0}
		q := []*TreeNode{s}
		for len(q) > 0 {
			u := q[0]
			q = q[1:]
			best = max(best, dist[u])
			for _, v := range adj[u] {
				if _, ok := dist[v]; !ok {
					dist[v] = dist[u] + 1
					q = append(q, v)
				}
			}
		}
	}
	return best
}

func TestExamples(t *testing.T) {
	if got := diameterOfBinaryTree(build([]int{1, 2, 3, 4, 5})); got != 3 {
		t.Errorf("got %d want 3", got)
	}
	if got := diameterOfBinaryTree(build([]int{1, 2})); got != 1 {
		t.Errorf("got %d want 1", got)
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(543, 1))
	for range 500 {
		root := randTree(r, 1+r.IntN(40), 0, 9)
		if got, want := diameterOfBinaryTree(root), brute(root); got != want {
			t.Fatalf("%v: got %d want %d", serialize(root), got, want)
		}
	}
}
