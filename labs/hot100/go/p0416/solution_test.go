package p0416

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
		{[]int{1}, false},
		{[]int{2, 2}, true},
	}
	for _, c := range cases {
		if got := canPartition(c.in); got != c.want {
			t.Errorf("canPartition(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}

// 暴力：枚举所有子集，看是否有和为总和一半的
func brute(nums []int) bool {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	for mask := range 1 << len(nums) {
		s := 0
		for i, x := range nums {
			if mask>>i&1 == 1 {
				s += x
			}
		}
		if 2*s == sum {
			return true
		}
	}
	return false
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(416, 2026))
	for range 1000 {
		nums := make([]int, 1+rng.IntN(10))
		for i := range nums {
			nums[i] = 1 + rng.IntN(12)
		}
		if got, want := canPartition(nums), brute(nums); got != want {
			t.Fatalf("canPartition(%v) = %v, want %v", nums, got, want)
		}
	}
}
