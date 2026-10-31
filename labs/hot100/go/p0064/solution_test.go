package p0064

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
		{[][]int{{5}}, 5},
	}
	for _, c := range cases {
		if got := minPathSum(c.grid); got != c.want {
			t.Errorf("minPathSum(%v) = %d, want %d", c.grid, got, c.want)
		}
	}
}

// 暴力递归：从 (i,j) 出发到右下角的最小和
func brute(g [][]int, i, j int) int {
	m, n := len(g), len(g[0])
	if i == m-1 && j == n-1 {
		return g[i][j]
	}
	best := -1
	if i+1 < m {
		best = brute(g, i+1, j)
	}
	if j+1 < n {
		if r := brute(g, i, j+1); best < 0 || r < best {
			best = r
		}
	}
	return g[i][j] + best
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(64, 2026))
	for range 500 {
		m, n := 1+rng.IntN(6), 1+rng.IntN(6)
		g := make([][]int, m)
		for i := range g {
			g[i] = make([]int, n)
			for j := range g[i] {
				g[i][j] = rng.IntN(10)
			}
		}
		if got, want := minPathSum(g), brute(g, 0, 0); got != want {
			t.Fatalf("minPathSum(%v) = %d, want %d", g, got, want)
		}
	}
}
