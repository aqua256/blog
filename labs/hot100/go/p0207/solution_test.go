package p0207

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	if !canFinish(2, [][]int{{1, 0}}) {
		t.Error("example 1 should be true")
	}
	if canFinish(2, [][]int{{1, 0}, {0, 1}}) {
		t.Error("example 2 should be false")
	}
}

// brute 传递闭包（Floyd-Warshall）：有某点能到达自己即有环。
func brute(n int, pre [][]int) bool {
	reach := make([][]bool, n)
	for i := range reach {
		reach[i] = make([]bool, n)
	}
	for _, p := range pre {
		reach[p[1]][p[0]] = true
	}
	for k := range n {
		for i := range n {
			for j := range n {
				if reach[i][k] && reach[k][j] {
					reach[i][j] = true
				}
			}
		}
	}
	for i := range n {
		if reach[i][i] {
			return false
		}
	}
	return true
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 207))
	for range 1000 {
		n := 1 + r.IntN(7)
		var pre [][]int
		for range r.IntN(n * 2) {
			a, b := r.IntN(n), r.IntN(n)
			if a != b {
				pre = append(pre, []int{a, b})
			}
		}
		if got, want := canFinish(n, pre), brute(n, pre); got != want {
			t.Fatalf("canFinish(%d,%v) = %v, want %v", n, pre, got, want)
		}
	}
}
