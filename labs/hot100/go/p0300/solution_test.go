package p0300

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
	}
	for _, c := range cases {
		if got := lengthOfLIS(c.in); got != c.want {
			t.Errorf("lengthOfLIS(%v) = %d, want %d", c.in, got, c.want)
		}
		if got := lengthOfLISDP(c.in); got != c.want {
			t.Errorf("lengthOfLISDP(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}

// 暴力：枚举所有子序列，检查严格递增
func brute(nums []int) int {
	best := 0
	for mask := range 1 << len(nums) {
		last, cnt, ok := 0, 0, true
		for i, x := range nums {
			if mask>>i&1 == 0 {
				continue
			}
			if cnt > 0 && x <= last {
				ok = false
				break
			}
			last, cnt = x, cnt+1
		}
		if ok {
			best = max(best, cnt)
		}
	}
	return best
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(300, 2026))
	for range 500 {
		nums := make([]int, 1+rng.IntN(12))
		for i := range nums {
			nums[i] = rng.IntN(10) - 3 // 小值域，制造重复和负数
		}
		want := brute(nums)
		if got := lengthOfLIS(nums); got != want {
			t.Fatalf("lengthOfLIS(%v) = %d, want %d", nums, got, want)
		}
		if got := lengthOfLISDP(nums); got != want {
			t.Fatalf("lengthOfLISDP(%v) = %d, want %d", nums, got, want)
		}
	}
}
