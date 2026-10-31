package p0114

import (
	"math/rand/v2"
	"slices"
	"testing"
)

// 展开后：所有 Left 为 nil，沿 Right 走得到的就是原树前序（节点身份也要一致）
func check(t *testing.T, root *TreeNode) {
	t.Helper()
	orig := serialize(root)
	want := nodes(root)
	flatten(root)
	var got []*TreeNode
	for n := root; n != nil; n = n.Right {
		if n.Left != nil {
			t.Fatalf("%v: left not nil at %d", orig, n.Val)
		}
		got = append(got, n)
	}
	if !slices.Equal(got, want) {
		t.Fatalf("%v: wrong order", orig)
	}
}

func TestExamples(t *testing.T) {
	root := build([]int{1, 2, 5, 3, 4, null, 6})
	flatten(root)
	want := []int{1, null, 2, null, 3, null, 4, null, 5, null, 6}
	if got := serialize(root); !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
	flatten(nil)
	check(t, build([]int{0}))
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(114, 1))
	for range 500 {
		check(t, randTree(r, r.IntN(40), 0, 9))
	}
}
