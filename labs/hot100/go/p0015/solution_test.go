package p0015

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"testing"
)

// 规范化：每个三元组内部排序，再整体排序成字符串列表
func normalize(ts [][]int) []string {
	out := make([]string, 0, len(ts))
	for _, t := range ts {
		t = slices.Clone(t)
		slices.Sort(t)
		out = append(out, fmt.Sprint(t))
	}
	slices.Sort(out)
	return out
}

func brute(nums []int) [][]int {
	set := map[[3]int]bool{}
	var ans [][]int
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					t := []int{nums[i], nums[j], nums[k]}
					slices.Sort(t)
					key := [3]int{t[0], t[1], t[2]}
					if !set[key] {
						set[key] = true
						ans = append(ans, t)
					}
				}
			}
		}
	}
	return ans
}

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, nil},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}
	for _, c := range cases {
		if got := threeSum(slices.Clone(c.in)); !slices.Equal(normalize(got), normalize(c.want)) {
			t.Fatalf("in=%v got=%v", c.in, got)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(15, 15))
	for it := 0; it < 500; it++ {
		nums := make([]int, 3+r.IntN(12))
		for i := range nums {
			nums[i] = r.IntN(11) - 5
		}
		want := normalize(brute(nums))
		got := normalize(threeSum(slices.Clone(nums)))
		// 规范化后完全相等 => 每个三元组合法、无重复、且数量与暴力一致
		if !slices.Equal(got, want) {
			t.Fatalf("nums=%v got=%v want=%v", nums, got, want)
		}
	}
}
