package p0560

import (
	"math/rand/v2"
	"testing"
)

func brute(nums []int, k int) int {
	ans := 0
	for i := range nums {
		s := 0
		for j := i; j < len(nums); j++ {
			s += nums[j]
			if s == k {
				ans++
			}
		}
	}
	return ans
}

func TestExamples(t *testing.T) {
	if got := subarraySum([]int{1, 1, 1}, 2); got != 2 {
		t.Fatalf("got %d", got)
	}
	if got := subarraySum([]int{1, 2, 3}, 3); got != 2 {
		t.Fatalf("got %d", got)
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(560, 560))
	for it := 0; it < 500; it++ {
		nums := make([]int, 1+r.IntN(15))
		for i := range nums {
			nums[i] = r.IntN(7) - 3 // 含负数和 0
		}
		k := r.IntN(9) - 4
		if got, want := subarraySum(nums, k), brute(nums, k); got != want {
			t.Fatalf("nums=%v k=%d got=%d want=%d", nums, k, got, want)
		}
	}
}
