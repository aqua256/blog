package p0005

import (
	"math/rand/v2"
	"strings"
	"testing"
)

func isPal(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// 暴力：枚举所有子串，取最长回文长度
func bruteLen(s string) int {
	best := 0
	for i := range len(s) {
		for j := i + 1; j <= len(s); j++ {
			if isPal(s[i:j]) {
				best = max(best, j-i)
			}
		}
	}
	return best
}

// 答案不唯一：检查是子串、是回文、长度最大
func check(t *testing.T, name, s, got string) {
	t.Helper()
	if !strings.Contains(s, got) || !isPal(got) || len(got) != bruteLen(s) {
		t.Fatalf("%s(%q) = %q, want a longest palindromic substring (len %d)", name, s, got, bruteLen(s))
	}
}

var funcs = map[string]func(string) string{
	"longestPalindrome":   longestPalindrome,
	"longestPalindromeDP": longestPalindromeDP,
}

func TestExamples(t *testing.T) {
	for name, f := range funcs {
		for _, s := range []string{"babad", "cbbd", "a", "ac", "forgeeksskeegfor"} {
			check(t, name, s, f(s))
		}
		if got := f("cbbd"); got != "bb" {
			t.Errorf("%s(cbbd) = %q, want bb", name, got)
		}
	}
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(5, 2026))
	for range 1000 {
		b := make([]byte, 1+rng.IntN(14))
		alpha := 1 + rng.IntN(3)
		for i := range b {
			b[i] = byte('a' + rng.IntN(alpha))
		}
		s := string(b)
		for name, f := range funcs {
			check(t, name, s, f(s))
		}
	}
}
