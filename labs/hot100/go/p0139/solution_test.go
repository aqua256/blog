package p0139

import (
	"math/rand/v2"
	"strings"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		s    string
		dict []string
		want bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}
	for _, c := range cases {
		if got := wordBreak(c.s, c.dict); got != c.want {
			t.Errorf("wordBreak(%q, %v) = %v, want %v", c.s, c.dict, got, c.want)
		}
	}
}

// 暴力递归：尝试每个单词作为前缀
func brute(s string, dict []string) bool {
	if s == "" {
		return true
	}
	for _, w := range dict {
		if strings.HasPrefix(s, w) && brute(s[len(w):], dict) {
			return true
		}
	}
	return false
}

func randStr(rng *rand.Rand, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte('a' + rng.IntN(2))
	}
	return string(b)
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(139, 2026))
	for range 500 {
		dict := make([]string, 1+rng.IntN(4))
		for i := range dict {
			dict[i] = randStr(rng, 1+rng.IntN(3))
		}
		s := randStr(rng, 1+rng.IntN(12))
		if got, want := wordBreak(s, dict), brute(s, dict); got != want {
			t.Fatalf("wordBreak(%q, %v) = %v, want %v", s, dict, got, want)
		}
	}
}
