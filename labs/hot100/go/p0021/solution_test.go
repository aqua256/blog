package p0021

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct{ a, b, want []int }{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{0}, []int{0}},
	}
	for _, c := range cases {
		if got := toSlice(mergeTwoLists(fromSlice(c.a), fromSlice(c.b))); !slices.Equal(got, c.want) {
			t.Errorf("merge(%v,%v) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}

func randSorted(r *rand.Rand) []int {
	s := make([]int, r.IntN(10))
	for i := range s {
		s[i] = r.IntN(20) - 10
	}
	slices.Sort(s)
	return s
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 21))
	for range 500 {
		a, b := randSorted(r), randSorted(r)
		want := slices.Concat(a, b)
		slices.Sort(want)
		if got := toSlice(mergeTwoLists(fromSlice(a), fromSlice(b))); !slices.Equal(got, want) {
			t.Fatalf("merge(%v,%v) = %v, want %v", a, b, got, want)
		}
	}
}
