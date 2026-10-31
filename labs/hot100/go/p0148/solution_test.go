package p0148

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct{ in, want []int }{
		{[]int{4, 2, 1, 3}, []int{1, 2, 3, 4}},
		{[]int{-1, 5, 3, 4, 0}, []int{-1, 0, 3, 4, 5}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		if got := toSlice(sortList(fromSlice(c.in))); !slices.Equal(got, c.want) {
			t.Errorf("sortList(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 148))
	for range 500 {
		in := make([]int, r.IntN(30))
		for i := range in {
			in[i] = r.IntN(21) - 10
		}
		want := slices.Clone(in)
		slices.Sort(want)
		if got := toSlice(sortList(fromSlice(in))); !slices.Equal(got, want) {
			t.Fatalf("sortList(%v) = %v, want %v", in, got, want)
		}
	}
}
