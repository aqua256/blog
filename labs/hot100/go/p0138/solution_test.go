package p0138

import (
	"math/rand/v2"
	"testing"
)

// check 验证深拷贝：结构一致、节点全是新的、random 指向对应新节点、原表未被改动。
func check(t *testing.T, spec [][2]int) {
	t.Helper()
	head, orig := build(spec)
	got := collect(copyRandomList(head))
	if len(got) != len(spec) {
		t.Fatalf("spec %v: copy has %d nodes, want %d", spec, len(got), len(spec))
	}
	isOld := map[*Node]bool{}
	for _, n := range orig {
		isOld[n] = true
	}
	index := map[*Node]int{}
	for i, n := range got {
		if isOld[n] {
			t.Fatalf("spec %v: node %d is an original node, not a copy", spec, i)
		}
		index[n] = i
	}
	for i, n := range got {
		if n.Val != spec[i][0] {
			t.Fatalf("spec %v: node %d val %d", spec, i, n.Val)
		}
		if spec[i][1] < 0 {
			if n.Random != nil {
				t.Fatalf("spec %v: node %d random should be nil", spec, i)
			}
			continue
		}
		j, ok := index[n.Random]
		if !ok || j != spec[i][1] {
			t.Fatalf("spec %v: node %d random points to wrong node", spec, i)
		}
	}
	// 原链表保持不变
	for i, n := range orig {
		var wantNext *Node
		if i+1 < len(orig) {
			wantNext = orig[i+1]
		}
		var wantRand *Node
		if spec[i][1] >= 0 {
			wantRand = orig[spec[i][1]]
		}
		if n.Next != wantNext || n.Random != wantRand || n.Val != spec[i][0] {
			t.Fatalf("spec %v: original list modified at %d", spec, i)
		}
	}
}

func TestExamples(t *testing.T) {
	check(t, [][2]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}})
	check(t, [][2]int{{1, 1}, {2, 1}})
	check(t, [][2]int{{3, -1}, {3, 0}, {3, -1}})
	check(t, nil)
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 138))
	for range 500 {
		n := r.IntN(10)
		spec := make([][2]int, n)
		for i := range spec {
			spec[i] = [2]int{r.IntN(3), r.IntN(n+1) - 1}
		}
		check(t, spec)
	}
}
