package p0200

import (
	"math/rand/v2"
	"testing"
)

func parse(rows []string) [][]byte {
	g := make([][]byte, len(rows))
	for i, r := range rows {
		g[i] = []byte(r)
	}
	return g
}

func TestExamples(t *testing.T) {
	cases := []struct {
		grid []string
		want int
	}{
		{[]string{"11110", "11010", "11000", "00000"}, 1},
		{[]string{"11000", "11000", "00100", "00011"}, 3},
	}
	for _, c := range cases {
		if got := numIslands(parse(c.grid)); got != c.want {
			t.Errorf("numIslands(%v) = %d, want %d", c.grid, got, c.want)
		}
	}
}

// brute 并查集计数，和 DFS 走完全不同的路子。
func brute(g [][]byte) int {
	m, n := len(g), len(g[0])
	parent := make([]int, m*n)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		for parent[x] != x {
			x = parent[x]
		}
		return x
	}
	for i := range m {
		for j := range n {
			if g[i][j] != '1' {
				continue
			}
			if i+1 < m && g[i+1][j] == '1' {
				parent[find(i*n+j)] = find((i+1)*n + j)
			}
			if j+1 < n && g[i][j+1] == '1' {
				parent[find(i*n+j)] = find(i*n + j + 1)
			}
		}
	}
	count := 0
	for i := range m {
		for j := range n {
			if g[i][j] == '1' && find(i*n+j) == i*n+j {
				count++
			}
		}
	}
	return count
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 200))
	for range 500 {
		m, n := 1+r.IntN(7), 1+r.IntN(7)
		g := make([][]byte, m)
		cp := make([][]byte, m)
		for i := range g {
			g[i] = make([]byte, n)
			for j := range g[i] {
				g[i][j] = "01"[r.IntN(2)]
			}
			cp[i] = append([]byte(nil), g[i]...)
		}
		want := brute(cp)
		if got := numIslands(g); got != want {
			t.Fatalf("numIslands(%q) = %d, want %d", cp, got, want)
		}
	}
}
