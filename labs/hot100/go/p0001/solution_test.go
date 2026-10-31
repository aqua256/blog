package p0001

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
	}{{[]int{2, 7, 11, 15}, 9}, {[]int{3, 2, 4}, 6}, {[]int{3, 3}, 6}}
	for _, c := range cases {
		checkValid(t, c.nums, c.target, twoSum(c.nums, c.target))
	}
}

func checkValid(t *testing.T, nums []int, target int, got []int) {
	t.Helper()
	if len(got) != 2 || got[0] == got[1] || got[0] < 0 || got[1] < 0 ||
		got[0] >= len(nums) || got[1] >= len(nums) || nums[got[0]]+nums[got[1]] != target {
		t.Fatalf("nums=%v target=%d got=%v", nums, target, got)
	}
}

func hasPair(nums []int, target int) bool {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return true
			}
		}
	}
	return false
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 1))
	for it := 0; it < 500; it++ {
		n := 2 + r.IntN(10)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = r.IntN(21) - 10
		}
		target := r.IntN(41) - 20
		got := twoSum(nums, target)
		if !hasPair(nums, target) {
			if got != nil {
				t.Fatalf("expected nil, nums=%v target=%d got=%v", nums, target, got)
			}
			continue
		}
		checkValid(t, nums, target, got)
	}
}
