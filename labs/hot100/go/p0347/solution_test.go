package p0347

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2, []int{1, 2}},
	}
	for _, c := range cases {
		got := topKFrequent(c.nums, c.k)
		slices.Sort(got)
		if !slices.Equal(got, c.want) {
			t.Errorf("topKFrequent(%v,%d) = %v, want %v", c.nums, c.k, got, c.want)
		}
	}
}

// valid 检查：k 个互不相同的数，且结果里的最低频次 >= 结果外的最高频次（允许并列）。
func valid(nums []int, k int, got []int) bool {
	freq := map[int]int{}
	for _, x := range nums {
		freq[x]++
	}
	if len(got) != k {
		return false
	}
	in := map[int]bool{}
	minIn := len(nums) + 1
	for _, x := range got {
		if in[x] || freq[x] == 0 {
			return false
		}
		in[x] = true
		minIn = min(minIn, freq[x])
	}
	for x, c := range freq {
		if !in[x] && c > minIn {
			return false
		}
	}
	return true
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 347))
	for range 1000 {
		nums := make([]int, 1+r.IntN(25))
		for i := range nums {
			nums[i] = r.IntN(8)
		}
		distinct := len(slices.Compact(slices.Sorted(slices.Values(nums))))
		k := 1 + r.IntN(distinct)
		if got := topKFrequent(nums, k); !valid(nums, k, got) {
			t.Fatalf("topKFrequent(%v,%d) = %v invalid", nums, k, got)
		}
	}
}
