package p0206

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct{ in, want []int }{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		if got := toSlice(reverseList(fromSlice(c.in))); !slices.Equal(got, c.want) {
			t.Errorf("reverseList(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 206))
	for range 500 {
		in := make([]int, r.IntN(12))
		for i := range in {
			in[i] = r.IntN(100)
		}
		want := slices.Clone(in)
		slices.Reverse(want)
		if got := toSlice(reverseList(fromSlice(in))); !slices.Equal(got, want) {
			t.Fatalf("reverseList(%v) = %v, want %v", in, got, want)
		}
	}
}
