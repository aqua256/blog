package p0023

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func run(in [][]int) []int {
	lists := make([]*ListNode, len(in))
	for i, s := range in {
		lists[i] = fromSlice(s)
	}
	return toSlice(mergeKLists(lists))
}

func TestExamples(t *testing.T) {
	cases := []struct {
		in   [][]int
		want []int
	}{
		{[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{[][]int{}, []int{}},
		{[][]int{{}}, []int{}},
	}
	for _, c := range cases {
		if got := run(c.in); !slices.Equal(got, c.want) {
			t.Errorf("mergeKLists(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 23))
	for range 500 {
		in := make([][]int, r.IntN(6))
		var want []int
		for i := range in {
			in[i] = make([]int, r.IntN(6))
			for j := range in[i] {
				in[i][j] = r.IntN(21) - 10
			}
			slices.Sort(in[i])
			want = append(want, in[i]...)
		}
		slices.Sort(want)
		if want == nil {
			want = []int{}
		}
		if got := run(in); !slices.Equal(got, want) {
			t.Fatalf("mergeKLists(%v) = %v, want %v", in, got, want)
		}
	}
}
