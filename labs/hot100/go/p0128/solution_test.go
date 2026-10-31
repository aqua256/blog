package p0128

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func brute(nums []int) int {
	s := slices.Clone(nums)
	slices.Sort(s)
	s = slices.Compact(s)
	best, cur := 0, 0
	for i := range s {
		if i > 0 && s[i] == s[i-1]+1 {
			cur++
		} else {
			cur = 1
		}
		best = max(best, cur)
	}
	return best
}

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{1, 0, 1, 2}, 3},
		{[]int{}, 0},
	}
	for _, c := range cases {
		if got := longestConsecutive(c.in); got != c.want {
			t.Fatalf("in=%v got=%d want=%d", c.in, got, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(128, 128))
	for it := 0; it < 500; it++ {
		nums := make([]int, r.IntN(15))
		for i := range nums {
			nums[i] = r.IntN(25) - 12
		}
		if got, want := longestConsecutive(nums), brute(nums); got != want {
			t.Fatalf("nums=%v got=%d want=%d", nums, got, want)
		}
	}
}
