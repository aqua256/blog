package p0033

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{1}, 1, 0},
		{[]int{3, 1}, 1, 1},
	}
	for _, c := range cases {
		if got := search(c.nums, c.target); got != c.want {
			t.Errorf("search(%v,%d)=%d want %d", c.nums, c.target, got, c.want)
		}
	}
}

func genRotated(r *rand.Rand, n int) []int {
	nums := make([]int, n)
	cur := r.IntN(10) - 20
	for i := range nums {
		cur += 1 + r.IntN(3)
		nums[i] = cur
	}
	k := r.IntN(n)
	return append(slices.Clone(nums[k:]), nums[:k]...)
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(33, 1))
	for range 3000 {
		nums := genRotated(r, 1+r.IntN(12))
		var target int
		if r.IntN(2) == 0 {
			target = nums[r.IntN(len(nums))]
		} else {
			target = r.IntN(60) - 25
		}
		want := slices.Index(nums, target)
		if got := search(nums, target); got != want {
			t.Fatalf("search(%v,%d)=%d want %d", nums, target, got, want)
		}
	}
}
