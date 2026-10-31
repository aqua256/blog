package p0153

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
		{[]int{2, 1}, 1},
		{[]int{1}, 1},
	}
	for _, c := range cases {
		if got := findMin(c.nums); got != c.want {
			t.Errorf("findMin(%v)=%d want %d", c.nums, got, c.want)
		}
	}
}

// genRotated 生成互不相同的升序数组并随机旋转
func genRotated(r *rand.Rand, n int) []int {
	nums := make([]int, n)
	cur := r.IntN(10) - 20
	for i := range nums {
		cur += 1 + r.IntN(3)
		nums[i] = cur
	}
	k := r.IntN(n)
	return append(slices.Clone(nums[k:]), nums[:k]...)
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(153, 1))
	for range 2000 {
		nums := genRotated(r, 1+r.IntN(12))
		if got, want := findMin(nums), slices.Min(nums); got != want {
			t.Fatalf("findMin(%v)=%d want %d", nums, got, want)
		}
	}
}
