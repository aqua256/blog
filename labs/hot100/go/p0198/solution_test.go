package p0198

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{5}, 5},
		{[]int{2, 1, 1, 2}, 4},
	}
	for _, c := range cases {
		if got := rob(c.in); got != c.want {
			t.Errorf("rob(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}

// 暴力：枚举所有不相邻子集
func brute(nums []int) int {
	best := 0
	for mask := range 1 << len(nums) {
		if mask&(mask>>1) != 0 {
			continue
		}
		sum := 0
		for i, x := range nums {
			if mask>>i&1 == 1 {
				sum += x
			}
		}
		best = max(best, sum)
	}
	return best
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(198, 2026))
	for range 500 {
		nums := make([]int, 1+rng.IntN(12))
		for i := range nums {
			nums[i] = rng.IntN(20)
		}
		if got, want := rob(nums), brute(nums); got != want {
			t.Fatalf("rob(%v) = %d, want %d", nums, got, want)
		}
	}
}
