package p0003

import (
	"math/rand/v2"
	"testing"
)

func brute(s string) int {
	best := 0
	for i := 0; i < len(s); i++ {
		seen := map[byte]bool{}
		for j := i; j < len(s) && !seen[s[j]]; j++ {
			seen[s[j]] = true
			best = max(best, j-i+1)
		}
	}
	return best
}

func TestExamples(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{{"abcabcbb", 3}, {"bbbbb", 1}, {"pwwkew", 3}, {"", 0}, {"abba", 2}}
	for _, c := range cases {
		if got := lengthOfLongestSubstring(c.s); got != c.want {
			t.Fatalf("s=%q got=%d want=%d", c.s, got, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(3, 3))
	for it := 0; it < 500; it++ {
		b := make([]byte, r.IntN(15))
		for i := range b {
			b[i] = 'a' + byte(r.IntN(4))
		}
		s := string(b)
		if got, want := lengthOfLongestSubstring(s), brute(s); got != want {
			t.Fatalf("s=%q got=%d want=%d", s, got, want)
		}
	}
}
