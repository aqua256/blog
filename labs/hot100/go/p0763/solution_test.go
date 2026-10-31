package p0763

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		in   string
		want []int
	}{
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
		{"eccbbbbdec", []int{10}},
		{"a", []int{1}},
	}
	for _, c := range cases {
		if got := partitionLabels(c.in); !slices.Equal(got, c.want) {
			t.Errorf("partitionLabels(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

// 合法：段长为正、总长等于 len(s)、每个字母只出现在一个段里
func valid(s string, parts []int) bool {
	seg := make(map[byte]int)
	pos := 0
	for k, p := range parts {
		if p <= 0 || pos+p > len(s) {
			return false
		}
		for i := pos; i < pos+p; i++ {
			if g, ok := seg[s[i]]; ok && g != k {
				return false
			}
			seg[s[i]] = k
		}
		pos += p
	}
	return pos == len(s)
}

// 暴力：枚举所有切分方式，取合法切分的最大段数
func bruteMax(s string) int {
	n := len(s)
	best := 0
	for mask := range 1 << (n - 1) {
		var parts []int
		prev := 0
		for i := 1; i < n; i++ {
			if mask>>(i-1)&1 == 1 {
				parts = append(parts, i-prev)
				prev = i
			}
		}
		parts = append(parts, n-prev)
		if valid(s, parts) {
			best = max(best, len(parts))
		}
	}
	return best
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(763, 2026))
	for range 500 {
		n := 1 + rng.IntN(12)
		alpha := 1 + rng.IntN(6)
		b := make([]byte, n)
		for i := range b {
			b[i] = byte('a' + rng.IntN(alpha))
		}
		s := string(b)
		got := partitionLabels(s)
		if !valid(s, got) {
			t.Fatalf("partitionLabels(%q) = %v is not a valid partition", s, got)
		}
		if want := bruteMax(s); len(got) != want {
			t.Fatalf("partitionLabels(%q) = %v, want %d parts", s, got, want)
		}
	}
}
