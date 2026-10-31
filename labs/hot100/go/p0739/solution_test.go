package p0739

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct{ in, want []int }{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
	}
	for _, c := range cases {
		if got := dailyTemperatures(c.in); !slices.Equal(got, c.want) {
			t.Errorf("dailyTemperatures(%v)=%v want %v", c.in, got, c.want)
		}
	}
}

func brute(temps []int) []int {
	ans := make([]int, len(temps))
	for i := range temps {
		for j := i + 1; j < len(temps); j++ {
			if temps[j] > temps[i] {
				ans[i] = j - i
				break
			}
		}
	}
	return ans
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(739, 1))
	for range 2000 {
		n := 1 + r.IntN(15)
		temps := make([]int, n)
		for i := range temps {
			temps[i] = 30 + r.IntN(6) // 小值域制造大量相等
		}
		if got, want := dailyTemperatures(temps), brute(temps); !slices.Equal(got, want) {
			t.Fatalf("dailyTemperatures(%v)=%v want %v", temps, got, want)
		}
	}
}
