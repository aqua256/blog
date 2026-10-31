package p0230

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	if got := kthSmallest(build([]int{3, 1, 4, null, 2}), 1); got != 1 {
		t.Errorf("got %d want 1", got)
	}
	if got := kthSmallest(build([]int{5, 3, 6, 2, 4, null, null, 1}), 3); got != 3 {
		t.Errorf("got %d want 3", got)
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(230, 1))
	for range 300 {
		root := randBST(r, 1+r.IntN(40))
		// 暴力：收集所有值排序
		var vals []int
		for _, n := range nodes(root) {
			vals = append(vals, n.Val)
		}
		slices.Sort(vals)
		for k := 1; k <= len(vals); k++ {
			if got := kthSmallest(root, k); got != vals[k-1] {
				t.Fatalf("%v k=%d: got %d want %d", serialize(root), k, got, vals[k-1])
			}
		}
	}
}
