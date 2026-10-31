package p0152

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2}, -2},
		{[]int{-2, 3, -4}, 24},
	}
	for _, c := range cases {
		if got := maxProduct(c.in); got != c.want {
			t.Errorf("maxProduct(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}

// 暴力：枚举所有子数组
func brute(nums []int) int {
	best := nums[0]
	for i := range nums {
		p := 1
		for j := i; j < len(nums); j++ {
			p *= nums[j]
			best = max(best, p)
		}
	}
	return best
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(152, 2026))
	for range 1000 {
		nums := make([]int, 1+rng.IntN(10))
		for i := range nums {
			nums[i] = rng.IntN(11) - 5
		}
		if got, want := maxProduct(nums), brute(nums); got != want {
			t.Fatalf("maxProduct(%v) = %d, want %d", nums, got, want)
		}
	}
}
