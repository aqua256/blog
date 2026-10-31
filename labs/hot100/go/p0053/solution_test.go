package p0053

import (
	"math/rand/v2"
	"testing"
)

func brute(nums []int) int {
	best := nums[0]
	for i := range nums {
		s := 0
		for j := i; j < len(nums); j++ {
			s += nums[j]
			best = max(best, s)
		}
	}
	return best
}

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6}, {[]int{1}, 1}, {[]int{5, 4, -1, 7, 8}, 23}, {[]int{-3, -1, -2}, -1}}
	for _, c := range cases {
		if got := maxSubArray(c.in); got != c.want {
			t.Fatalf("in=%v got=%d want=%d", c.in, got, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(53, 53))
	for it := 0; it < 500; it++ {
		nums := make([]int, 1+r.IntN(15))
		for i := range nums {
			nums[i] = r.IntN(21) - 10
		}
		if got, want := maxSubArray(nums), brute(nums); got != want {
			t.Fatalf("nums=%v got=%d want=%d", nums, got, want)
		}
	}
}
