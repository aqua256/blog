package p0074

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	m := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	if !searchMatrix(m, 3) {
		t.Error("want true for 3")
	}
	if searchMatrix(m, 13) {
		t.Error("want false for 13")
	}
	if searchMatrix([][]int{{1}}, 2) || !searchMatrix([][]int{{1}}, 1) {
		t.Error("single cell")
	}
}

func brute(matrix [][]int, target int) bool {
	for _, row := range matrix {
		for _, x := range row {
			if x == target {
				return true
			}
		}
	}
	return false
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(74, 1))
	for range 2000 {
		m, n := 1+r.IntN(5), 1+r.IntN(5)
		matrix := make([][]int, m)
		cur := r.IntN(10) - 20
		for i := range matrix {
			matrix[i] = make([]int, n)
			for j := range matrix[i] {
				cur += 1 + r.IntN(3)
				matrix[i][j] = cur
			}
		}
		target := r.IntN(100) - 25
		if got, want := searchMatrix(matrix, target), brute(matrix, target); got != want {
			t.Fatalf("searchMatrix(%v,%d)=%v want %v", matrix, target, got, want)
		}
	}
}
