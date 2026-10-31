package tpl_dp

import (
	"math/rand/v2"
	"testing"
)

func bruteLinear(nums []int) int {
	best := 0
	for mask := range 1 << len(nums) {
		if mask&(mask>>1) != 0 {
			continue
		}
		s := 0
		for i, x := range nums {
			if mask>>i&1 == 1 {
				s += x
			}
		}
		best = max(best, s)
	}
	return best
}

func brute01(w, v []int, c int) int {
	best := 0
	for mask := range 1 << len(w) {
		tw, tv := 0, 0
		for i := range w {
			if mask>>i&1 == 1 {
				tw, tv = tw+w[i], tv+v[i]
			}
		}
		if tw <= c {
			best = max(best, tv)
		}
	}
	return best
}

func bruteComplete(w []int, t int) int {
	if t == 0 {
		return 0
	}
	best := -1
	for _, x := range w {
		if x <= t {
			if r := bruteComplete(w, t-x); r >= 0 && (best < 0 || r+1 < best) {
				best = r + 1
			}
		}
	}
	return best
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(1, 2026))
	for range 300 {
		n := 1 + rng.IntN(10)
		w, v := make([]int, n), make([]int, n)
		for i := range n {
			w[i], v[i] = 1+rng.IntN(8), rng.IntN(20)
		}
		if got, want := linear(v), bruteLinear(v); got != want {
			t.Fatalf("linear(%v) = %d, want %d", v, got, want)
		}
		c := rng.IntN(25)
		if got, want := knapsack01(w, v, c), brute01(w, v, c); got != want {
			t.Fatalf("knapsack01(%v, %v, %d) = %d, want %d", w, v, c, got, want)
		}
		ws := w[:min(n, 3)]
		if got, want := knapsackCompleteMin(ws, c), bruteComplete(ws, c); got != want {
			t.Fatalf("knapsackCompleteMin(%v, %d) = %d, want %d", ws, c, got, want)
		}
	}
}
