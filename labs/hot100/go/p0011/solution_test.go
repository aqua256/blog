package p0011

import (
	"math/rand/v2"
	"testing"
)

func brute(h []int) int {
	best := 0
	for i := range h {
		for j := i + 1; j < len(h); j++ {
			best = max(best, min(h[i], h[j])*(j-i))
		}
	}
	return best
}

func TestExamples(t *testing.T) {
	if got := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}); got != 49 {
		t.Fatalf("got %d", got)
	}
	if got := maxArea([]int{1, 1}); got != 1 {
		t.Fatalf("got %d", got)
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(11, 11))
	for it := 0; it < 500; it++ {
		h := make([]int, 2+r.IntN(12))
		for i := range h {
			h[i] = r.IntN(10)
		}
		if got, want := maxArea(h), brute(h); got != want {
			t.Fatalf("h=%v got=%d want=%d", h, got, want)
		}
	}
}
