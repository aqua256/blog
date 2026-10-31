package p0072

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"", "abc", 3},
		{"abc", "", 3},
		{"", "", 0},
	}
	for _, c := range cases {
		if got := minDistance(c.a, c.b); got != c.want {
			t.Errorf("minDistance(%q, %q) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

// 暴力递归（无记忆化，指数级），从末尾比较
func brute(a, b string) int {
	if a == "" {
		return len(b)
	}
	if b == "" {
		return len(a)
	}
	if a[len(a)-1] == b[len(b)-1] {
		return brute(a[:len(a)-1], b[:len(b)-1])
	}
	return 1 + min(
		brute(a[:len(a)-1], b[:len(b)-1]), // 替换
		brute(a[:len(a)-1], b),            // 删除
		brute(a, b[:len(b)-1]),            // 插入
	)
}

func randStr(rng *rand.Rand, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte('a' + rng.IntN(3))
	}
	return string(b)
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(72, 2026))
	for range 500 {
		a, b := randStr(rng, rng.IntN(7)), randStr(rng, rng.IntN(7))
		if got, want := minDistance(a, b), brute(a, b); got != want {
			t.Fatalf("minDistance(%q, %q) = %d, want %d", a, b, got, want)
		}
	}
}
