package p0042

import (
	"math/rand/v2"
	"testing"
)

// 暴力：每个位置的水位 = min(左侧最高, 右侧最高) - 自身高度
func brute(h []int) int {
	total := 0
	for i := range h {
		lm, rm := 0, 0
		for j := 0; j <= i; j++ {
			lm = max(lm, h[j])
		}
		for j := i; j < len(h); j++ {
			rm = max(rm, h[j])
		}
		total += min(lm, rm) - h[i]
	}
	return total
}

func TestExamples(t *testing.T) {
	if got := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}); got != 6 {
		t.Fatalf("got %d", got)
	}
	if got := trap([]int{4, 2, 0, 3, 2, 5}); got != 9 {
		t.Fatalf("got %d", got)
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(42, 42))
	for it := 0; it < 500; it++ {
		h := make([]int, 1+r.IntN(15))
		for i := range h {
			h[i] = r.IntN(8)
		}
		if got, want := trap(h), brute(h); got != want {
			t.Fatalf("h=%v got=%d want=%d", h, got, want)
		}
	}
}
