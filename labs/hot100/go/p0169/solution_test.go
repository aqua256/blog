package p0169

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1}, 1},
	}
	for _, c := range cases {
		if got := majorityElement(c.in); got != c.want {
			t.Errorf("majorityElement(%v)=%d want %d", c.in, got, c.want)
		}
	}
}

func brute(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
		if cnt[x] > len(nums)/2 {
			return x
		}
	}
	panic("no majority")
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(169, 1))
	for range 2000 {
		n := 1 + r.IntN(15)
		major := r.IntN(5)
		k := n/2 + 1 + r.IntN(n-n/2) // 众数出现次数 > n/2
		nums := make([]int, 0, n)
		for range k {
			nums = append(nums, major)
		}
		for len(nums) < n {
			nums = append(nums, r.IntN(5))
		}
		r.Shuffle(n, func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
		if got, want := majorityElement(nums), brute(nums); got != want {
			t.Fatalf("majorityElement(%v)=%d want %d", nums, got, want)
		}
	}
}
