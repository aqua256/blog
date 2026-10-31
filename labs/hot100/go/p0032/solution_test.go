package p0032

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := map[string]int{
		"(()":    2,
		")()())": 4,
		"":       0,
		"()(())": 6,
		"(()())": 6,
		"())()":  2,
	}
	for s, want := range cases {
		if got := longestValidParentheses(s); got != want {
			t.Errorf("longestValidParentheses(%q) = %d, want %d", s, got, want)
		}
	}
}

func isValid(s string) bool {
	bal := 0
	for i := range len(s) {
		if s[i] == '(' {
			bal++
		} else if bal--; bal < 0 {
			return false
		}
	}
	return bal == 0
}

// 暴力：枚举所有子串逐个检查
func brute(s string) int {
	best := 0
	for i := range len(s) {
		for j := i + 2; j <= len(s); j += 2 {
			if isValid(s[i:j]) {
				best = max(best, j-i)
			}
		}
	}
	return best
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(32, 2026))
	for range 2000 {
		b := make([]byte, rng.IntN(16))
		for i := range b {
			b[i] = "()"[rng.IntN(2)]
		}
		s := string(b)
		if got, want := longestValidParentheses(s), brute(s); got != want {
			t.Fatalf("longestValidParentheses(%q) = %d, want %d", s, got, want)
		}
	}
}
