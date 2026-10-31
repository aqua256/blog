package p0239

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func brute(nums []int, k int) []int {
	var ans []int
	for i := 0; i+k <= len(nums); i++ {
		ans = append(ans, slices.Max(nums[i:i+k]))
	}
	return ans
}

func TestExamples(t *testing.T) {
	if got := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3); !slices.Equal(got, []int{3, 3, 5, 5, 6, 7}) {
		t.Fatalf("got %v", got)
	}
	if got := maxSlidingWindow([]int{1}, 1); !slices.Equal(got, []int{1}) {
		t.Fatalf("got %v", got)
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(239, 239))
	for it := 0; it < 500; it++ {
		nums := make([]int, 1+r.IntN(15))
		for i := range nums {
			nums[i] = r.IntN(9) - 4
		}
		k := 1 + r.IntN(len(nums))
		if got, want := maxSlidingWindow(nums, k), brute(nums, k); !slices.Equal(got, want) {
			t.Fatalf("nums=%v k=%d got=%v want=%v", nums, k, got, want)
		}
	}
}
