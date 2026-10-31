package p0056

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

// 暴力：在数轴上标记覆盖。用 2x 坐标，使 [1,2] 和 [3,4] 之间的空隙也能被区分
func brute(intervals [][]int) [][]int {
	var cov [64]bool
	for _, in := range intervals {
		for x := 2 * in[0]; x <= 2*in[1]; x++ {
			cov[x] = true
		}
	}
	var ans [][]int
	for x := 0; x < len(cov); x++ {
		if !cov[x] {
			continue
		}
		y := x
		for y+1 < len(cov) && cov[y+1] {
			y++
		}
		ans = append(ans, []int{x / 2, y / 2})
		x = y
	}
	return ans
}

func clone(a [][]int) [][]int {
	out := make([][]int, len(a))
	for i := range a {
		out[i] = []int{a[i][0], a[i][1]}
	}
	return out
}

func TestExamples(t *testing.T) {
	cases := []struct{ in, want [][]int }{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{4, 7}, {1, 4}}, [][]int{{1, 7}}},
	}
	for _, c := range cases {
		if got := merge(clone(c.in)); fmt.Sprint(got) != fmt.Sprint(c.want) {
			t.Fatalf("in=%v got=%v", c.in, got)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(56, 56))
	for it := 0; it < 500; it++ {
		in := make([][]int, 1+r.IntN(8))
		for i := range in {
			a := r.IntN(20)
			in[i] = []int{a, a + r.IntN(5)}
		}
		if got, want := merge(clone(in)), brute(in); fmt.Sprint(got) != fmt.Sprint(want) {
			t.Fatalf("in=%v got=%v want=%v", in, got, want)
		}
	}
}
