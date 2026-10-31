package p0105

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func pre(n *TreeNode, out *[]int) {
	if n != nil {
		*out = append(*out, n.Val)
		pre(n.Left, out)
		pre(n.Right, out)
	}
}

func in(n *TreeNode, out *[]int) {
	if n != nil {
		in(n.Left, out)
		*out = append(*out, n.Val)
		in(n.Right, out)
	}
}

func TestExamples(t *testing.T) {
	got := serialize(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
	if want := []int{3, 9, 20, null, null, 15, 7}; !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
	if got := serialize(buildTree([]int{-1}, []int{-1})); !slices.Equal(got, []int{-1}) {
		t.Errorf("got %v", got)
	}
}

// 随机树 -> 前序+中序 -> 重建 -> 序列化必须等于原树
func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(105, 1))
	for range 500 {
		root := randDistinctTree(r, 1+r.IntN(40))
		var p, i []int
		pre(root, &p)
		in(root, &i)
		if got, want := serialize(buildTree(p, i)), serialize(root); !slices.Equal(got, want) {
			t.Fatalf("pre=%v in=%v: got %v want %v", p, i, got, want)
		}
	}
}
