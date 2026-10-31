package p0094

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func bruteInorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := bruteInorder(root.Left)
	res = append(res, root.Val)
	return append(res, bruteInorder(root.Right)...)
}

func TestExamples(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{1, null, 2, 3}, []int{1, 3, 2}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4, 5, null, 8, null, null, 6, 7, 9}, []int{4, 2, 6, 5, 7, 1, 3, 9, 8}},
	}
	for _, c := range cases {
		if got := inorderTraversal(build(c.in)); !slices.Equal(got, c.want) {
			t.Errorf("%v: got %v want %v", c.in, got, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(94, 1))
	for range 500 {
		root := randTree(r, r.IntN(30), -50, 50)
		if got, want := inorderTraversal(root), bruteInorder(root); !slices.Equal(got, want) {
			t.Fatalf("%v: got %v want %v", serialize(root), got, want)
		}
	}
}
