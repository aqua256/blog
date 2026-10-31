package p0031

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct{ in, want []int }{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1}, []int{1}},
		{[]int{1, 5, 1}, []int{5, 1, 1}},
	}
	for _, c := range cases {
		nextPermutation(c.in)
		if !slices.Equal(c.in, c.want) {
			t.Errorf("got %v want %v", c.in, c.want)
		}
	}
}

// permutations 枚举全排列（含重复元素产生的重复排列）
func permutations(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{slices.Clone(nums)}
	}
	var res [][]int
	for i := range nums {
		rest := slices.Concat(nums[:i], nums[i+1:])
		for _, p := range permutations(rest) {
			res = append(res, append([]int{nums[i]}, p...))
		}
	}
	return res
}

// brute：全排列排序去重，取严格大于 nums 的第一个；没有则取最小的
func brute(nums []int) []int {
	perms := permutations(nums)
	slices.SortFunc(perms, slices.Compare)
	perms = slices.CompactFunc(perms, slices.Equal)
	for _, p := range perms {
		if slices.Compare(p, nums) > 0 {
			return p
		}
	}
	return perms[0]
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(31, 1))
	for range 1500 {
		n := 1 + r.IntN(6)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = r.IntN(4) // 小值域制造重复
		}
		want := brute(nums)
		orig := slices.Clone(nums)
		nextPermutation(nums)
		if !slices.Equal(nums, want) {
			t.Fatalf("nextPermutation(%v)=%v want %v", orig, nums, want)
		}
	}
}
