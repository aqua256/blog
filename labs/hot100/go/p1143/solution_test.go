package p1143

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"bsbininm", "jmjkbkjkv", 1},
	}
	for _, c := range cases {
		if got := longestCommonSubsequence(c.a, c.b); got != c.want {
			t.Errorf("LCS(%q, %q) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

// 暴力递归（无记忆化，指数级）
func brute(a, b string) int {
	if a == "" || b == "" {
		return 0
	}
	if a[0] == b[0] {
		return 1 + brute(a[1:], b[1:])
	}
	return max(brute(a[1:], b), brute(a, b[1:]))
}

func randStr(rng *rand.Rand, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte('a' + rng.IntN(3))
	}
	return string(b)
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(1143, 2026))
	for range 500 {
		a, b := randStr(rng, 1+rng.IntN(8)), randStr(rng, 1+rng.IntN(8))
		if got, want := longestCommonSubsequence(a, b), brute(a, b); got != want {
			t.Fatalf("LCS(%q, %q) = %d, want %d", a, b, got, want)
		}
	}
}
