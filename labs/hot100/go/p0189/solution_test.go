package p0189

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func brute(nums []int, k int) []int {
	n := len(nums)
	out := make([]int, n)
	for i, x := range nums {
		out[(i+k)%n] = x
	}
	return out
}

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}
	for _, c := range cases {
		got := slices.Clone(c.in)
		rotate(got, c.k)
		if !slices.Equal(got, c.want) {
			t.Fatalf("in=%v got=%v", c.in, got)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(189, 189))
	for it := 0; it < 500; it++ {
		nums := make([]int, 1+r.IntN(10))
		for i := range nums {
			nums[i] = r.IntN(100)
		}
		k := r.IntN(25)
		got := slices.Clone(nums)
		rotate(got, k)
		if want := brute(nums, k); !slices.Equal(got, want) {
			t.Fatalf("nums=%v k=%d got=%v want=%v", nums, k, got, want)
		}
	}
}
