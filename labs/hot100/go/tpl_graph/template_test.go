package tpl_graph

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestFloodFill(t *testing.T) {
	g := [][]byte{[]byte("110"), []byte("010"), []byte("001")}
	floodFill(g, 0, 0)
	if string(g[0]) != "000" || string(g[1]) != "000" || string(g[2]) != "001" {
		t.Fatalf("got %q", g)
	}
}

func TestBFSLevels(t *testing.T) {
	g := [][]int{{1, 1, 1}, {0, 0, 1}, {1, 1, 1}}
	got := bfsLevels(g, [][2]int{{0, 0}})
	want := [][]int{{0, 1, 2}, {-1, -1, 3}, {6, 5, 4}}
	for i := range want {
		if !slices.Equal(got[i], want[i]) {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestTopoSortRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 9))
	for range 500 {
		n := 1 + r.IntN(8)
		// 随机排列 perm，只加 perm 里从前往后的边 -> 必定无环
		perm := r.Perm(n)
		var edges [][2]int
		for range r.IntN(2 * n) {
			a, b := r.IntN(n), r.IntN(n)
			if a < b {
				edges = append(edges, [2]int{perm[a], perm[b]})
			}
		}
		order := topoSort(n, edges)
		if len(order) != n {
			t.Fatalf("DAG not fully sorted: %v", order)
		}
		pos := make([]int, n)
		for i, v := range order {
			pos[v] = i
		}
		for _, e := range edges {
			if pos[e[0]] >= pos[e[1]] {
				t.Fatalf("edge %v violated in %v", e, order)
			}
		}
		// 加一条反向边成环（若存在边）
		if len(edges) > 0 {
			e := edges[0]
			if len(topoSort(n, append(edges, [2]int{e[1], e[0]}))) == n {
				t.Fatalf("cycle not detected")
			}
		}
	}
}
