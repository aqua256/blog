package p0041

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func brute(nums []int) int {
	for x := 1; ; x++ {
		if !slices.Contains(nums, x) {
			return x
		}
	}
}

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{{[]int{1, 2, 0}, 3}, {[]int{3, 4, -1, 1}, 2}, {[]int{7, 8, 9, 11, 12}, 1}, {[]int{1, 1}, 2}}
	for _, c := range cases {
		if got := firstMissingPositive(slices.Clone(c.in)); got != c.want {
			t.Fatalf("in=%v got=%d want=%d", c.in, got, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(41, 41))
	for it := 0; it < 500; it++ {
		nums := make([]int, 1+r.IntN(12))
		for i := range nums {
			nums[i] = r.IntN(18) - 4 // 含负数、0、重复、越界大数
		}
		if got, want := firstMissingPositive(slices.Clone(nums)), brute(nums); got != want {
			t.Fatalf("nums=%v got=%d want=%d", nums, got, want)
		}
	}
}
