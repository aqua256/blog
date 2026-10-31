package p0238

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func brute(nums []int) []int {
	out := make([]int, len(nums))
	for i := range nums {
		p := 1
		for j, x := range nums {
			if j != i {
				p *= x
			}
		}
		out[i] = p
	}
	return out
}

func TestExamples(t *testing.T) {
	if got := productExceptSelf([]int{1, 2, 3, 4}); !slices.Equal(got, []int{24, 12, 8, 6}) {
		t.Fatalf("got %v", got)
	}
	if got := productExceptSelf([]int{-1, 1, 0, -3, 3}); !slices.Equal(got, []int{0, 0, 9, 0, 0}) {
		t.Fatalf("got %v", got)
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(238, 238))
	for it := 0; it < 500; it++ {
		nums := make([]int, 2+r.IntN(8))
		for i := range nums {
			nums[i] = r.IntN(7) - 3 // 含 0，覆盖一个/多个 0 的情况
		}
		if got, want := productExceptSelf(nums), brute(nums); !slices.Equal(got, want) {
			t.Fatalf("nums=%v got=%v want=%v", nums, got, want)
		}
	}
}
