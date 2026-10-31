package p0034

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
		{[]int{1}, 1, []int{0, 0}},
	}
	for _, c := range cases {
		if got := searchRange(c.nums, c.target); !slices.Equal(got, c.want) {
			t.Errorf("searchRange(%v,%d)=%v want %v", c.nums, c.target, got, c.want)
		}
	}
}

func brute(nums []int, target int) []int {
	res := []int{-1, -1}
	for i, x := range nums {
		if x == target {
			if res[0] == -1 {
				res[0] = i
			}
			res[1] = i
		}
	}
	return res
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(34, 1))
	for range 2000 {
		n := r.IntN(12)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = r.IntN(7) - 3
		}
		slices.Sort(nums)
		target := r.IntN(9) - 4
		if got, want := searchRange(nums, target), brute(nums, target); !slices.Equal(got, want) {
			t.Fatalf("searchRange(%v,%d)=%v want %v", nums, target, got, want)
		}
	}
}
