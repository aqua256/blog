package p0075

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct{ in, want []int }{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{0}, []int{0}},
	}
	for _, c := range cases {
		sortColors(c.in)
		if !slices.Equal(c.in, c.want) {
			t.Errorf("got %v want %v", c.in, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(75, 1))
	for range 3000 {
		n := 1 + r.IntN(15)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = r.IntN(3)
		}
		want := slices.Clone(nums)
		slices.Sort(want)
		orig := slices.Clone(nums)
		sortColors(nums)
		if !slices.Equal(nums, want) {
			t.Fatalf("sortColors(%v)=%v want %v", orig, nums, want)
		}
	}
}
