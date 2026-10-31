package p0121

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1}, 0},
	}
	for _, c := range cases {
		if got := maxProfit(c.in); got != c.want {
			t.Errorf("maxProfit(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}

func brute(prices []int) int {
	best := 0
	for i := range prices {
		for j := i + 1; j < len(prices); j++ {
			best = max(best, prices[j]-prices[i])
		}
	}
	return best
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(121, 2026))
	for range 500 {
		n := 1 + rng.IntN(12)
		prices := make([]int, n)
		for i := range prices {
			prices[i] = rng.IntN(20)
		}
		if got, want := maxProfit(prices), brute(prices); got != want {
			t.Fatalf("maxProfit(%v) = %d, want %d", prices, got, want)
		}
	}
}
