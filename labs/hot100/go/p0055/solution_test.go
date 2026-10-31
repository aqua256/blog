package p0055

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{0}, true},
		{[]int{0, 1}, false},
	}
	for _, c := range cases {
		if got := canJump(c.in); got != c.want {
			t.Errorf("canJump(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}

// 暴力：reach[i] = 存在 j<i 可达且 j+nums[j] >= i
func brute(nums []int) bool {
	reach := make([]bool, len(nums))
	reach[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if reach[j] && j+nums[j] >= i {
				reach[i] = true
				break
			}
		}
	}
	return reach[len(nums)-1]
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(55, 2026))
	for range 1000 {
		n := 1 + rng.IntN(12)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rng.IntN(4)
		}
		if got, want := canJump(nums), brute(nums); got != want {
			t.Fatalf("canJump(%v) = %v, want %v", nums, got, want)
		}
	}
}
