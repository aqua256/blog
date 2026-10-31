package p0136

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
		{[]int{-3, 5, 5}, -3},
	}
	for _, c := range cases {
		if got := singleNumber(c.in); got != c.want {
			t.Errorf("singleNumber(%v)=%d want %d", c.in, got, c.want)
		}
	}
}

func brute(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	for x, c := range cnt {
		if c == 1 {
			return x
		}
	}
	panic("no single")
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(136, 1))
	for range 1000 {
		pairs := r.IntN(6)
		vals := r.Perm(20)[:pairs+1] // 互不相同
		nums := []int{vals[0] - 10}
		for _, v := range vals[1:] {
			nums = append(nums, v-10, v-10)
		}
		r.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
		if got, want := singleNumber(nums), brute(nums); got != want {
			t.Fatalf("singleNumber(%v)=%d want %d", nums, got, want)
		}
	}
}
