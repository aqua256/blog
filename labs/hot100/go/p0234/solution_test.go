package p0234

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	if !isPalindrome(fromSlice([]int{1, 2, 2, 1})) {
		t.Error("[1,2,2,1] should be palindrome")
	}
	if isPalindrome(fromSlice([]int{1, 2})) {
		t.Error("[1,2] should not be palindrome")
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 234))
	for range 1000 {
		n := 1 + r.IntN(10)
		in := make([]int, n)
		for i := range in {
			in[i] = r.IntN(3)
		}
		if r.IntN(2) == 0 { // 一半强制构造回文
			for i := range n / 2 {
				in[n-1-i] = in[i]
			}
		}
		rev := slices.Clone(in)
		slices.Reverse(rev)
		if got, want := isPalindrome(fromSlice(in)), slices.Equal(in, rev); got != want {
			t.Fatalf("isPalindrome(%v) = %v, want %v", in, got, want)
		}
	}
}
