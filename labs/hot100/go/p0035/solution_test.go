package p0035

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1}, 1, 0},
	}
	for _, c := range cases {
		if got := searchInsert(c.nums, c.target); got != c.want {
			t.Errorf("searchInsert(%v,%d)=%d want %d", c.nums, c.target, got, c.want)
		}
	}
}

func brute(nums []int, target int) int {
	for i, x := range nums {
		if x >= target {
			return i
		}
	}
	return len(nums)
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(35, 1))
	for range 1000 {
		n := 1 + r.IntN(12)
		nums := make([]int, n)
		cur := r.IntN(5) - 10
		for i := range nums { // 严格递增（题目保证无重复）
			cur += 1 + r.IntN(3)
			nums[i] = cur
		}
		target := r.IntN(40) - 15
		if got, want := searchInsert(nums, target), brute(nums, target); got != want {
			t.Fatalf("searchInsert(%v,%d)=%d want %d", nums, target, got, want)
		}
	}
}
