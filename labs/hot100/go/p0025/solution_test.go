package p0025

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{[]int{1, 2, 3}, 1, []int{1, 2, 3}},
	}
	for _, c := range cases {
		if got := toSlice(reverseKGroup(fromSlice(c.in), c.k)); !slices.Equal(got, c.want) {
			t.Errorf("reverseKGroup(%v,%d) = %v, want %v", c.in, c.k, got, c.want)
		}
	}
}

func brute(in []int, k int) []int {
	out := slices.Clone(in)
	for i := 0; i+k <= len(out); i += k {
		slices.Reverse(out[i : i+k])
	}
	return out
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 25))
	for range 500 {
		n := 1 + r.IntN(15)
		in := make([]int, n)
		for i := range in {
			in[i] = i
		}
		k := 1 + r.IntN(n)
		if got, want := toSlice(reverseKGroup(fromSlice(in), k)), brute(in, k); !slices.Equal(got, want) {
			t.Fatalf("reverseKGroup(%v,%d) = %v, want %v", in, k, got, want)
		}
	}
}
