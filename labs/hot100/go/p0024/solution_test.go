package p0024

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct{ in, want []int }{
		{[]int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3}, []int{2, 1, 3}},
	}
	for _, c := range cases {
		if got := toSlice(swapPairs(fromSlice(c.in))); !slices.Equal(got, c.want) {
			t.Errorf("swapPairs(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 24))
	for range 500 {
		in := make([]int, r.IntN(12))
		for i := range in {
			in[i] = i
		}
		want := slices.Clone(in)
		for i := 0; i+1 < len(want); i += 2 {
			want[i], want[i+1] = want[i+1], want[i]
		}
		if got := toSlice(swapPairs(fromSlice(in))); !slices.Equal(got, want) {
			t.Fatalf("swapPairs(%v) = %v, want %v", in, got, want)
		}
	}
}
