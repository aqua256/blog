package tpl_sliding_window

import (
	"math/rand/v2"
	"testing"
)

func TestLongestValid(t *testing.T) {
	r := rand.New(rand.NewPCG(7, 7))
	for it := 0; it < 300; it++ {
		b := make([]byte, r.IntN(12))
		for i := range b {
			b[i] = 'a' + byte(r.IntN(3))
		}
		k := 1 + r.IntN(2)
		want := 0
		for i := range b {
			cnt := map[byte]int{}
			for j := i; j < len(b); j++ {
				cnt[b[j]]++
				if cnt[b[j]] > k {
					break
				}
				want = max(want, j-i+1)
			}
		}
		if got := longestValid(string(b), k); got != want {
			t.Fatalf("s=%q k=%d got=%d want=%d", b, k, got, want)
		}
	}
}

func TestShortestValid(t *testing.T) {
	r := rand.New(rand.NewPCG(8, 8))
	for it := 0; it < 300; it++ {
		nums := make([]int, 1+r.IntN(10))
		for i := range nums {
			nums[i] = 1 + r.IntN(5)
		}
		target := 1 + r.IntN(20)
		want := 0
		for i := range nums {
			s := 0
			for j := i; j < len(nums); j++ {
				s += nums[j]
				if s >= target {
					if want == 0 || j-i+1 < want {
						want = j - i + 1
					}
					break
				}
			}
		}
		if got := shortestValid(nums, target); got != want {
			t.Fatalf("nums=%v target=%d got=%d want=%d", nums, target, got, want)
		}
	}
}

func TestMaxSumFixed(t *testing.T) {
	r := rand.New(rand.NewPCG(9, 9))
	for it := 0; it < 300; it++ {
		nums := make([]int, 1+r.IntN(10))
		for i := range nums {
			nums[i] = r.IntN(11) - 5
		}
		k := 1 + r.IntN(len(nums))
		want := -1 << 60
		for i := 0; i+k <= len(nums); i++ {
			s := 0
			for _, x := range nums[i : i+k] {
				s += x
			}
			want = max(want, s)
		}
		if got := maxSumFixed(nums, k); got != want {
			t.Fatalf("nums=%v k=%d got=%d want=%d", nums, k, got, want)
		}
	}
}
