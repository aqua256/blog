package p0287

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{3, 3, 3, 3, 3}, 3},
		{[]int{1, 1}, 1},
		{[]int{2, 2, 2}, 2},
	}
	for _, c := range cases {
		orig := slices.Clone(c.in)
		if got := findDuplicate(c.in); got != c.want {
			t.Errorf("findDuplicate(%v)=%d want %d", orig, got, c.want)
		}
		if !slices.Equal(c.in, orig) {
			t.Errorf("input modified: %v -> %v", orig, c.in)
		}
	}
}

// O(n^2) 暴力，不修改数组
func brute(nums []int) int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	panic("no duplicate")
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(287, 1))
	for range 3000 {
		n := 1 + r.IntN(10) // 长度 n+1，值域 [1,n]
		dup := 1 + r.IntN(n)
		// 先放一个 1..n 的排列再补一个 dup，然后随机把若干其他值也替换成 dup
		nums := append(r.Perm(n), 0)
		for i := range n {
			nums[i]++
		}
		nums[n] = dup
		for i := range nums {
			if nums[i] != dup && r.IntN(4) == 0 {
				nums[i] = dup
			}
		}
		r.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
		orig := slices.Clone(nums)
		want := brute(nums)
		if got := findDuplicate(nums); got != want {
			t.Fatalf("findDuplicate(%v)=%d want %d", orig, got, want)
		}
		if !slices.Equal(nums, orig) {
			t.Fatalf("input modified: %v -> %v", orig, nums)
		}
	}
}
