package p0045

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{0}, 0},
		{[]int{1, 2}, 1},
	}
	for _, c := range cases {
		if got := jump(c.in); got != c.want {
			t.Errorf("jump(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}

// 对拍：dp[i] = 到达 i 的最少步数，-1 表示不可达
func dpJump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = -1
		for j := 0; j < i; j++ {
			if dp[j] >= 0 && j+nums[j] >= i && (dp[i] < 0 || dp[j]+1 < dp[i]) {
				dp[i] = dp[j] + 1
			}
		}
	}
	return dp[n-1]
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(45, 2026))
	tested := 0
	for tested < 500 {
		n := 1 + rng.IntN(14)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rng.IntN(5)
		}
		want := dpJump(nums)
		if want < 0 {
			continue // 题目保证可达
		}
		tested++
		if got := jump(nums); got != want {
			t.Fatalf("jump(%v) = %d, want %d", nums, got, want)
		}
	}
}
