package p0102

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

// 暴力：前序 DFS 带深度，按深度归桶（同层内自左向右）
func brute(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(*TreeNode, int)
	dfs = func(n *TreeNode, d int) {
		if n == nil {
			return
		}
		if d == len(res) {
			res = append(res, []int{})
		}
		res[d] = append(res[d], n.Val)
		dfs(n.Left, d+1)
		dfs(n.Right, d+1)
	}
	dfs(root, 0)
	return res
}

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want string
	}{
		{[]int{3, 9, 20, null, null, 15, 7}, "[[3] [9 20] [15 7]]"},
		{[]int{1}, "[[1]]"},
		{[]int{}, "[]"},
	}
	for _, c := range cases {
		if got := fmt.Sprint(levelOrder(build(c.in))); got != c.want {
			t.Errorf("%v: got %s want %s", c.in, got, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(102, 1))
	for range 500 {
		root := randTree(r, r.IntN(40), 0, 99)
		if got, want := fmt.Sprint(levelOrder(root)), fmt.Sprint(brute(root)); got != want {
			t.Fatalf("%v: got %s want %s", serialize(root), got, want)
		}
	}
}
