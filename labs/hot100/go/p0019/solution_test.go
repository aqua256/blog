package p0019

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		n    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
	}
	for _, c := range cases {
		if got := toSlice(removeNthFromEnd(fromSlice(c.in), c.n)); !slices.Equal(got, c.want) {
			t.Errorf("remove(%v,%d) = %v, want %v", c.in, c.n, got, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 19))
	for range 500 {
		sz := 1 + r.IntN(12)
		in := make([]int, sz)
		for i := range in {
			in[i] = i
		}
		n := 1 + r.IntN(sz)
		want := slices.Delete(slices.Clone(in), sz-n, sz-n+1)
		if got := toSlice(removeNthFromEnd(fromSlice(in), n)); !slices.Equal(got, want) {
			t.Fatalf("remove(%v,%d) = %v, want %v", in, n, got, want)
		}
	}
}
