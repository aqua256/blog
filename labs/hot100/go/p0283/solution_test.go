package p0283

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func brute(nums []int) []int {
	out := []int{}
	for _, x := range nums {
		if x != 0 {
			out = append(out, x)
		}
	}
	for len(out) < len(nums) {
		out = append(out, 0)
	}
	return out
}

func TestExamples(t *testing.T) {
	cases := []struct{ in, want []int }{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
	}
	for _, c := range cases {
		got := slices.Clone(c.in)
		moveZeroes(got)
		if !slices.Equal(got, c.want) {
			t.Fatalf("in=%v got=%v", c.in, got)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(283, 283))
	for it := 0; it < 500; it++ {
		nums := make([]int, r.IntN(12))
		for i := range nums {
			if r.IntN(2) == 0 {
				nums[i] = r.IntN(9) - 4
			}
		}
		got := slices.Clone(nums)
		moveZeroes(got)
		if want := brute(nums); !slices.Equal(got, want) {
			t.Fatalf("nums=%v got=%v want=%v", nums, got, want)
		}
	}
}
