package p0226

import (
	"math/rand/v2"
	"slices"
	"testing"
)

// 暴力：新建一棵镜像树，不修改原树
func mirror(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{Val: root.Val, Left: mirror(root.Right), Right: mirror(root.Left)}
}

func TestExamples(t *testing.T) {
	cases := []struct{ in, want []int }{
		{[]int{4, 2, 7, 1, 3, 6, 9}, []int{4, 7, 2, 9, 6, 3, 1}},
		{[]int{2, 1, 3}, []int{2, 3, 1}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		if got := serialize(invertTree(build(c.in))); !slices.Equal(got, c.want) {
			t.Errorf("%v: got %v want %v", c.in, got, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(226, 1))
	for range 500 {
		root := randTree(r, r.IntN(30), 0, 99)
		want := serialize(mirror(root))
		orig := serialize(root)
		got := serialize(invertTree(root))
		if !slices.Equal(got, want) {
			t.Fatalf("%v: got %v want %v", orig, got, want)
		}
		// 再翻一次回到原树
		if back := serialize(invertTree(root)); !slices.Equal(back, orig) {
			t.Fatalf("double invert %v != %v", back, orig)
		}
	}
}
