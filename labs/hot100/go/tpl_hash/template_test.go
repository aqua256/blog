package tpl_hash

import "testing"

func TestPairIndex(t *testing.T) {
	i, j, ok := pairIndex([]int{3, 2, 4}, 6)
	if !ok || i != 1 || j != 2 {
		t.Fatalf("got %d %d %v", i, j, ok)
	}
	if _, _, ok := pairIndex([]int{3}, 6); ok {
		t.Fatal("should not pair with itself")
	}
}

func TestGroupBy(t *testing.T) {
	g := groupBy([]int{1, 2, 3, 4, 5}, func(x int) int { return x % 2 })
	if len(g[0]) != 2 || len(g[1]) != 3 {
		t.Fatalf("got %v", g)
	}
}
