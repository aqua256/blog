package tpl_heap

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestTopKRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 13))
	for range 500 {
		nums := make([]int, 1+r.IntN(20))
		for i := range nums {
			nums[i] = r.IntN(21) - 10
		}
		k := 1 + r.IntN(len(nums))
		got := slices.Sorted(slices.Values(topK(nums, k)))
		want := slices.Sorted(slices.Values(nums))[len(nums)-k:]
		if !slices.Equal(got, want) {
			t.Fatalf("topK(%v,%d) = %v, want %v", nums, k, got, want)
		}
	}
}
