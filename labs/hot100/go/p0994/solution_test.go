package p0994

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
		{[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1},
		{[][]int{{0, 2}}, 0},
		{[][]int{{0}}, 0},
	}
	for _, c := range cases {
		if got := orangesRotting(c.grid); got != c.want {
			t.Errorf("got %d, want %d", got, c.want)
		}
	}
}

// brute 逐分钟整盘模拟，直到不再变化。
func brute(g [][]int) int {
	m, n := len(g), len(g[0])
	minutes := 0
	for {
		var rot [][2]int
		for i := range m {
			for j := range n {
				if g[i][j] != 1 {
					continue
				}
				if (i > 0 && g[i-1][j] == 2) || (i+1 < m && g[i+1][j] == 2) ||
					(j > 0 && g[i][j-1] == 2) || (j+1 < n && g[i][j+1] == 2) {
					rot = append(rot, [2]int{i, j})
				}
			}
		}
		if len(rot) == 0 {
			break
		}
		for _, p := range rot {
			g[p[0]][p[1]] = 2
		}
		minutes++
	}
	for i := range m {
		for j := range n {
			if g[i][j] == 1 {
				return -1
			}
		}
	}
	return minutes
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 994))
	for range 500 {
		m, n := 1+r.IntN(6), 1+r.IntN(6)
		g := make([][]int, m)
		cp := make([][]int, m)
		for i := range g {
			g[i] = make([]int, n)
			cp[i] = make([]int, n)
			for j := range g[i] {
				v := r.IntN(10)
				switch {
				case v < 2:
					g[i][j] = 0
				case v < 9:
					g[i][j] = 1
				default:
					g[i][j] = 2
				}
				cp[i][j] = g[i][j]
			}
		}
		orig := make([][]int, m)
		for i := range cp {
			orig[i] = append([]int(nil), cp[i]...)
		}
		want := brute(cp)
		if got := orangesRotting(g); got != want {
			t.Fatalf("orangesRotting(%v) = %d, want %d", orig, got, want)
		}
	}
}
