package p0108

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func inorder(n *TreeNode, out *[]int) {
	if n == nil {
		return
	}
	inorder(n.Left, out)
	*out = append(*out, n.Val)
	inorder(n.Right, out)
}

// 返回高度；不平衡返回 -1
func balancedHeight(n *TreeNode) int {
	if n == nil {
		return 0
	}
	l, r := balancedHeight(n.Left), balancedHeight(n.Right)
	if l < 0 || r < 0 || l-r > 1 || r-l > 1 {
		return -1
	}
	return max(l, r) + 1
}

// 答案不唯一：检查中序 == nums 且高度平衡
func check(t *testing.T, nums []int) {
	t.Helper()
	root := sortedArrayToBST(nums)
	got := []int{}
	inorder(root, &got)
	if !slices.Equal(got, nums) {
		t.Fatalf("%v: inorder %v", nums, got)
	}
	if balancedHeight(root) < 0 {
		t.Fatalf("%v: not balanced %v", nums, serialize(root))
	}
}

func TestExamples(t *testing.T) {
	check(t, []int{-10, -3, 0, 5, 9})
	check(t, []int{1, 3})
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(108, 1))
	for range 500 {
		n := r.IntN(60)
		nums := r.Perm(3*n + 1)[:n]
		slices.Sort(nums)
		check(t, nums)
	}
}
