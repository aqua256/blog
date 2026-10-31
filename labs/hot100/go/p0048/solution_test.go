package p0048

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

// 暴力：新矩阵 out[j][n-1-i] = a[i][j]
func brute(a [][]int) [][]int {
	n := len(a)
	out := make([][]int, n)
	for i := range out {
		out[i] = make([]int, n)
	}
	for i := range n {
		for j := range n {
			out[j][n-1-i] = a[i][j]
		}
	}
	return out
}

func clone(a [][]int) [][]int {
	out := make([][]int, len(a))
	for i := range a {
		out[i] = append([]int(nil), a[i]...)
	}
	return out
}

func TestExamples(t *testing.T) {
	got := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(got)
	if fmt.Sprint(got) != fmt.Sprint([][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}) {
		t.Fatalf("got %v", got)
	}
	got = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(got)
	if fmt.Sprint(got) != fmt.Sprint([][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}) {
		t.Fatalf("got %v", got)
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(48, 48))
	for it := 0; it < 300; it++ {
		n := 1 + r.IntN(6)
		a := make([][]int, n)
		for i := range a {
			a[i] = make([]int, n)
			for j := range a[i] {
				a[i][j] = r.IntN(100)
			}
		}
		got := clone(a)
		rotate(got)
		if want := brute(a); fmt.Sprint(got) != fmt.Sprint(want) {
			t.Fatalf("in=%v got=%v want=%v", a, got, want)
		}
	}
}
