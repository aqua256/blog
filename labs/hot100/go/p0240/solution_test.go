package p0240

import (
	"math/rand/v2"
	"testing"
)

func brute(a [][]int, target int) bool {
	for _, row := range a {
		for _, v := range row {
			if v == target {
				return true
			}
		}
	}
	return false
}

func TestExamples(t *testing.T) {
	a := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	if !searchMatrix(a, 5) {
		t.Fatal("5 should be found")
	}
	if searchMatrix(a, 20) {
		t.Fatal("20 should not be found")
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(240, 240))
	for it := 0; it < 500; it++ {
		m, n := 1+r.IntN(6), 1+r.IntN(6)
		// 生成行列均升序的矩阵：a[i][j] = max(上, 左) + 随机增量
		a := make([][]int, m)
		for i := range a {
			a[i] = make([]int, n)
			for j := range a[i] {
				base := 0
				if i > 0 {
					base = a[i-1][j]
				}
				if j > 0 {
					base = max(base, a[i][j-1])
				}
				a[i][j] = base + r.IntN(3)
			}
		}
		for target := -1; target <= a[m-1][n-1]+1; target++ {
			if got, want := searchMatrix(a, target), brute(a, target); got != want {
				t.Fatalf("a=%v target=%d got=%v want=%v", a, target, got, want)
			}
		}
	}
}
