package p0076

import (
	"math/rand/v2"
	"testing"
)

func covers(w, t string) bool {
	var c [128]int
	for i := range len(w) {
		c[w[i]]++
	}
	for i := range len(t) {
		c[t[i]]--
		if c[t[i]] < 0 {
			return false
		}
	}
	return true
}

// 暴力：最短覆盖长度（答案不唯一时只比较长度 + 合法性）
func bruteLen(s, t string) int {
	for l := 1; l <= len(s); l++ {
		for i := 0; i+l <= len(s); i++ {
			if covers(s[i:i+l], t) {
				return l
			}
		}
	}
	return 0
}

func TestExamples(t *testing.T) {
	cases := []struct{ s, t, want string }{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
	}
	for _, c := range cases {
		if got := minWindow(c.s, c.t); got != c.want {
			t.Fatalf("s=%q t=%q got=%q want=%q", c.s, c.t, got, c.want)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(76, 76))
	gen := func(n int) string {
		b := make([]byte, n)
		for i := range b {
			b[i] = "abcA"[r.IntN(4)]
		}
		return string(b)
	}
	for it := 0; it < 500; it++ {
		s, tt := gen(1+r.IntN(14)), gen(1+r.IntN(4))
		got, wantLen := minWindow(s, tt), bruteLen(s, tt)
		if len(got) != wantLen || (wantLen > 0 && !covers(got, tt)) {
			t.Fatalf("s=%q t=%q got=%q wantLen=%d", s, tt, got, wantLen)
		}
	}
}
