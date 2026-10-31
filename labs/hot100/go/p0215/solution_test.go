package p0215

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	if got := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2); got != 5 {
		t.Errorf("example 1 = %d, want 5", got)
	}
	if got := findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4); got != 4 {
		t.Errorf("example 2 = %d, want 4", got)
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 215))
	for range 1000 {
		nums := make([]int, 1+r.IntN(20))
		for i := range nums {
			nums[i] = r.IntN(21) - 10
		}
		k := 1 + r.IntN(len(nums))
		sorted := slices.Clone(nums)
		slices.Sort(sorted)
		want := sorted[len(sorted)-k]
		if got := findKthLargest(slices.Clone(nums), k); got != want {
			t.Fatalf("findKthLargest(%v,%d) = %d, want %d", nums, k, got, want)
		}
	}
}
