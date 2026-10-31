package p0295

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExample(t *testing.T) {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	if got := mf.FindMedian(); got != 1.5 {
		t.Fatalf("median = %v, want 1.5", got)
	}
	mf.AddNum(3)
	if got := mf.FindMedian(); got != 2 {
		t.Fatalf("median = %v, want 2", got)
	}
}

func brute(nums []int) float64 {
	s := slices.Sorted(slices.Values(nums))
	n := len(s)
	if n%2 == 1 {
		return float64(s[n/2])
	}
	return float64(s[n/2-1]+s[n/2]) / 2
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 295))
	for range 300 {
		mf := Constructor()
		var nums []int
		for range 1 + r.IntN(40) {
			x := r.IntN(41) - 20
			mf.AddNum(x)
			nums = append(nums, x)
			if got, want := mf.FindMedian(), brute(nums); got != want {
				t.Fatalf("after %v median = %v, want %v", nums, got, want)
			}
		}
	}
}
