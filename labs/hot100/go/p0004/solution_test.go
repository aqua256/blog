package p0004

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		a, b []int
		want float64
	}{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{}, []int{1}, 1},
		{[]int{2}, []int{}, 2},
		{[]int{1, 1}, []int{1, 1}, 1},
	}
	for _, c := range cases {
		if got := findMedianSortedArrays(c.a, c.b); got != c.want {
			t.Errorf("findMedian(%v,%v)=%v want %v", c.a, c.b, got, c.want)
		}
	}
}

func brute(a, b []int) float64 {
	all := slices.Concat(a, b)
	slices.Sort(all)
	n := len(all)
	if n%2 == 1 {
		return float64(all[n/2])
	}
	return float64(all[n/2-1]+all[n/2]) / 2
}

func gen(r *rand.Rand, n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = r.IntN(11) - 5
	}
	slices.Sort(a)
	return a
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(4, 1))
	for range 5000 {
		a, b := gen(r, r.IntN(8)), gen(r, r.IntN(8))
		if len(a)+len(b) == 0 {
			continue
		}
		if got, want := findMedianSortedArrays(a, b), brute(a, b); got != want {
			t.Fatalf("findMedian(%v,%v)=%v want %v", a, b, got, want)
		}
	}
}
