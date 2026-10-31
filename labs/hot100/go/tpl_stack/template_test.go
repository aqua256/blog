package tpl_stack

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(12, 1))
	for range 1000 {
		n := r.IntN(12)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = r.IntN(6)
		}
		ng, ps := make([]int, n), make([]int, n)
		for i := range nums {
			ng[i], ps[i] = -1, -1
			for j := i + 1; j < n; j++ {
				if nums[j] > nums[i] {
					ng[i] = j
					break
				}
			}
			for j := i - 1; j >= 0; j-- {
				if nums[j] < nums[i] {
					ps[i] = j
					break
				}
			}
		}
		if got := nextGreater(nums); !slices.Equal(got, ng) {
			t.Fatalf("nextGreater(%v)=%v want %v", nums, got, ng)
		}
		if got := prevSmaller(nums); !slices.Equal(got, ps) {
			t.Fatalf("prevSmaller(%v)=%v want %v", nums, got, ps)
		}
	}
}
