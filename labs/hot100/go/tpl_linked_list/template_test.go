package tpl_linked_list

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestTemplate(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 7))
	for range 500 {
		in := make([]int, r.IntN(12))
		for i := range in {
			in[i] = r.IntN(4)
		}
		v := r.IntN(4)
		want := slices.DeleteFunc(slices.Clone(in), func(x int) bool { return x == v })
		if got := toSlice(removeVal(fromSlice(in), v)); !slices.Equal(got, want) {
			t.Fatalf("removeVal(%v,%d) = %v, want %v", in, v, got, want)
		}
		rev := slices.Clone(in)
		slices.Reverse(rev)
		if got := toSlice(reverse(fromSlice(in))); !slices.Equal(got, rev) {
			t.Fatalf("reverse(%v) = %v", in, got)
		}
		if len(in) > 0 {
			if got := toSlice(middle(fromSlice(in))); !slices.Equal(got, in[len(in)/2:]) {
				t.Fatalf("middle(%v) = %v", in, got)
			}
		}
	}
}
