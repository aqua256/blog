package p0438

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func brute(s, p string) []int {
	sp := []byte(p)
	slices.Sort(sp)
	var ans []int
	for i := 0; i+len(p) <= len(s); i++ {
		w := []byte(s[i : i+len(p)])
		slices.Sort(w)
		if slices.Equal(w, sp) {
			ans = append(ans, i)
		}
	}
	return ans
}

func TestExamples(t *testing.T) {
	if got := findAnagrams("cbaebabacd", "abc"); !slices.Equal(got, []int{0, 6}) {
		t.Fatalf("got %v", got)
	}
	if got := findAnagrams("abab", "ab"); !slices.Equal(got, []int{0, 1, 2}) {
		t.Fatalf("got %v", got)
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(438, 438))
	gen := func(n int) string {
		b := make([]byte, n)
		for i := range b {
			b[i] = 'a' + byte(r.IntN(3))
		}
		return string(b)
	}
	for it := 0; it < 500; it++ {
		s, p := gen(1+r.IntN(15)), gen(1+r.IntN(4))
		if got, want := findAnagrams(s, p), brute(s, p); !slices.Equal(got, want) {
			t.Fatalf("s=%q p=%q got=%v want=%v", s, p, got, want)
		}
	}
}
