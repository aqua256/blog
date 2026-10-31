package p0073

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func clone(a [][]int) [][]int {
	out := make([][]int, len(a))
	for i := range a {
		out[i] = append([]int(nil), a[i]...)
	}
	return out
}

func brute(a [][]int) [][]int {
	out := clone(a)
	for i := range a {
		for j := range a[i] {
			if a[i][j] == 0 {
				for k := range a[i] {
					out[i][k] = 0
				}
				for k := range a {
					out[k][j] = 0
				}
			}
		}
	}
	return out
}

func TestExamples(t *testing.T) {
	cases := []struct{ in, want [][]int }{
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
	}
	for _, c := range cases {
		got := clone(c.in)
		setZeroes(got)
		if fmt.Sprint(got) != fmt.Sprint(c.want) {
			t.Fatalf("in=%v got=%v", c.in, got)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(73, 73))
	for it := 0; it < 500; it++ {
		m, n := 1+r.IntN(5), 1+r.IntN(5)
		a := make([][]int, m)
		for i := range a {
			a[i] = make([]int, n)
			for j := range a[i] {
				if r.IntN(5) != 0 {
					a[i][j] = 1 + r.IntN(9)
				}
			}
		}
		got := clone(a)
		setZeroes(got)
		if want := brute(a); fmt.Sprint(got) != fmt.Sprint(want) {
			t.Fatalf("in=%v got=%v want=%v", a, got, want)
		}
	}
}
