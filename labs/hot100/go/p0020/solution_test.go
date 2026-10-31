package p0020

import (
	"math/rand/v2"
	"strings"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := map[string]bool{
		"()": true, "()[]{}": true, "(]": false, "([])": true, "([)]": false,
		"{[]}": true, "(": false, "]": false, "((": false,
	}
	for s, want := range cases {
		if got := isValid(s); got != want {
			t.Errorf("isValid(%q)=%v want %v", s, got, want)
		}
	}
}

// 暴力：反复删掉相邻配对，最后为空即合法
func brute(s string) bool {
	for {
		t := strings.NewReplacer("()", "", "[]", "", "{}", "").Replace(s)
		if t == s {
			return s == ""
		}
		s = t
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(20, 1))
	const chars = "()[]{}"
	for range 5000 {
		n := 1 + r.IntN(10)
		b := make([]byte, n)
		for i := range b {
			b[i] = chars[r.IntN(len(chars))]
		}
		s := string(b)
		if got, want := isValid(s), brute(s); got != want {
			t.Fatalf("isValid(%q)=%v want %v", s, got, want)
		}
	}
}
