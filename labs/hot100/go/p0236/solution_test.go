package p0236

import (
	"math/rand/v2"
	"testing"
)

// 暴力：建父指针，p 的所有祖先入集合，q 向上第一个命中的就是 LCA
func brute(root, p, q *TreeNode) *TreeNode {
	parent := map[*TreeNode]*TreeNode{}
	for _, n := range nodes(root) {
		for _, c := range []*TreeNode{n.Left, n.Right} {
			if c != nil {
				parent[c] = n
			}
		}
	}
	anc := map[*TreeNode]bool{}
	for n := p; n != nil; n = parent[n] {
		anc[n] = true
	}
	for n := q; ; n = parent[n] {
		if anc[n] {
			return n
		}
	}
}

func find(root *TreeNode, v int) *TreeNode {
	for _, n := range nodes(root) {
		if n.Val == v {
			return n
		}
	}
	return nil
}

func TestExamples(t *testing.T) {
	root := build([]int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4})
	cases := [][3]int{{5, 1, 3}, {5, 4, 5}}
	for _, c := range cases {
		if got := lowestCommonAncestor(root, find(root, c[0]), find(root, c[1])); got.Val != c[2] {
			t.Errorf("p=%d q=%d: got %d want %d", c[0], c[1], got.Val, c[2])
		}
	}
	root = build([]int{1, 2})
	if got := lowestCommonAncestor(root, find(root, 1), find(root, 2)); got.Val != 1 {
		t.Errorf("got %d want 1", got.Val)
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(236, 1))
	for range 500 {
		root := randDistinctTree(r, 2+r.IntN(40))
		ns := nodes(root)
		for range 5 {
			p, q := ns[r.IntN(len(ns))], ns[r.IntN(len(ns))]
			if p == q {
				continue
			}
			if got, want := lowestCommonAncestor(root, p, q), brute(root, p, q); got != want {
				t.Fatalf("%v p=%d q=%d: got %d want %d", serialize(root), p.Val, q.Val, got.Val, want.Val)
			}
		}
	}
}
