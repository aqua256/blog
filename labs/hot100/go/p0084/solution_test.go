package p0084

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 4}, 4},
		{[]int{0}, 0},
		{[]int{2, 2, 2}, 6},
	}
	for _, c := range cases {
		if got := largestRectangleArea(c.in); got != c.want {
			t.Errorf("largestRectangleArea(%v)=%d want %d", c.in, got, c.want)
		}
	}
}

// O(n^2)：枚举左端点，向右扩展时维护区间最小高度
func brute(h []int) int {
	best := 0
	for i := range h {
		low := h[i]
		for j := i; j < len(h); j++ {
			low = min(low, h[j])
			best = max(best, low*(j-i+1))
		}
	}
	return best
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(84, 1))
	for range 3000 {
		n := 1 + r.IntN(15)
		h := make([]int, n)
		for i := range h {
			h[i] = r.IntN(7)
		}
		want := brute(h)
		if got := largestRectangleArea(h); got != want {
			t.Fatalf("largestRectangleArea(%v)=%d want %d", h, got, want)
		}
	}
}
