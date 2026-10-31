package p0054

import (
	"math/rand/v2"
	"slices"
	"testing"
)

// 暴力：模拟走格子，撞墙或已访问就右转
func brute(a [][]int) []int {
	m, n := len(a), len(a[0])
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var out []int
	i, j, d := 0, 0, 0
	for range m * n {
		out = append(out, a[i][j])
		seen[i][j] = true
		ni, nj := i+dirs[d][0], j+dirs[d][1]
		if ni < 0 || ni >= m || nj < 0 || nj >= n || seen[ni][nj] {
			d = (d + 1) % 4
			ni, nj = i+dirs[d][0], j+dirs[d][1]
		}
		i, j = ni, nj
	}
	return out
}

func TestExamples(t *testing.T) {
	if got := spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}); !slices.Equal(got, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}) {
		t.Fatalf("got %v", got)
	}
	if got := spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}); !slices.Equal(got, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}) {
		t.Fatalf("got %v", got)
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(54, 54))
	for it := 0; it < 300; it++ {
		m, n := 1+r.IntN(6), 1+r.IntN(6)
		a := make([][]int, m)
		for i := range a {
			a[i] = make([]int, n)
			for j := range a[i] {
				a[i][j] = i*n + j // 值互不相同，顺序错一点都能发现
			}
		}
		if got, want := spiralOrder(a), brute(a); !slices.Equal(got, want) {
			t.Fatalf("m=%d n=%d got=%v want=%v", m, n, got, want)
		}
	}
}
