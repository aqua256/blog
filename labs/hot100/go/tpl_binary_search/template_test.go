package tpl_binary_search

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestBounds(t *testing.T) {
	r := rand.New(rand.NewPCG(11, 1))
	for range 2000 {
		n := r.IntN(10)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = r.IntN(8)
		}
		slices.Sort(nums)
		x := r.IntN(10) - 1
		first, gt := n, n
		for i := n - 1; i >= 0; i-- {
			if nums[i] >= x {
				first = i
			}
			if nums[i] > x {
				gt = i
			}
		}
		if got := lowerBound(nums, x); got != first {
			t.Fatalf("lowerBound(%v,%d)=%d want %d", nums, x, got, first)
		}
		if got := upperBound(nums, x); got != gt {
			t.Fatalf("upperBound(%v,%d)=%d want %d", nums, x, got, gt)
		}
		if got := lastLess(nums, x); got != first-1 {
			t.Fatalf("lastLess(%v,%d)=%d", nums, x, got)
		}
		if got := lastLessEq(nums, x); got != gt-1 {
			t.Fatalf("lastLessEq(%v,%d)=%d", nums, x, got)
		}
		if got := firstTrue(0, n, func(i int) bool { return nums[i] >= x }); got != first {
			t.Fatalf("firstTrue mismatch")
		}
	}
}
