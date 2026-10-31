package p0101

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func mirrorCopy(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{Val: root.Val, Left: mirrorCopy(root.Right), Right: mirrorCopy(root.Left)}
}

// 暴力：对称 <=> 镜像后层序序列化不变
func brute(root *TreeNode) bool {
	return slices.Equal(serialize(root), serialize(mirrorCopy(root)))
}

func TestExamples(t *testing.T) {
	if !isSymmetric(build([]int{1, 2, 2, 3, 4, 4, 3})) {
		t.Error("example 1 should be true")
	}
	if isSymmetric(build([]int{1, 2, 2, null, 3, null, 3})) {
		t.Error("example 2 should be false")
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(101, 1))
	trues := 0
	for i := range 1000 {
		var root *TreeNode
		if i%2 == 0 {
			// 构造对称树，偶尔扰动一个值
			half := randTree(r, r.IntN(8), 1, 2)
			root = &TreeNode{Val: 0, Left: half, Right: mirrorCopy(half)}
			if r.IntN(3) == 0 {
				ns := nodes(root)
				ns[r.IntN(len(ns))].Val += 1
			}
		} else {
			root = randTree(r, 1+r.IntN(7), 1, 2)
		}
		want := brute(root)
		if want {
			trues++
		}
		if got := isSymmetric(root); got != want {
			t.Fatalf("%v: got %v want %v", serialize(root), got, want)
		}
	}
	if trues < 100 {
		t.Fatalf("too few symmetric cases: %d", trues)
	}
}
