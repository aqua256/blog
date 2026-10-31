package tpl_dp_2d

import (
	"math/rand/v2"
	"testing"
)

func bruteGrid(g [][]int, i, j int) int {
	m, n := len(g), len(g[0])
	if i == m-1 && j == n-1 {
		return g[i][j]
	}
	best := -1
	if i+1 < m {
		best = bruteGrid(g, i+1, j)
	}
	if j+1 < n {
		if r := bruteGrid(g, i, j+1); best < 0 || r < best {
			best = r
		}
	}
	return g[i][j] + best
}

func bruteLCS(a, b string) int {
	if a == "" || b == "" {
		return 0
	}
	if a[0] == b[0] {
		return 1 + bruteLCS(a[1:], b[1:])
	}
	return max(bruteLCS(a[1:], b), bruteLCS(a, b[1:]))
}

func brutePal(s string) int {
	count := 0
	for i := range len(s) {
		for j := i + 1; j <= len(s); j++ {
			ok := true
			for l, r := i, j-1; l < r; l, r = l+1, r-1 {
				ok = ok && s[l] == s[r]
			}
			if ok {
				count++
			}
		}
	}
	return count
}

func randStr(rng *rand.Rand, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte('a' + rng.IntN(3))
	}
	return string(b)
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(2, 2026))
	for range 300 {
		m, n := 1+rng.IntN(5), 1+rng.IntN(5)
		g := make([][]int, m)
		for i := range g {
			g[i] = make([]int, n)
			for j := range g[i] {
				g[i][j] = rng.IntN(10)
			}
		}
		if got, want := gridMinPath(g), bruteGrid(g, 0, 0); got != want {
			t.Fatalf("gridMinPath(%v) = %d, want %d", g, got, want)
		}
		a, b := randStr(rng, rng.IntN(8)), randStr(rng, rng.IntN(8))
		if got, want := twoStrings(a, b), bruteLCS(a, b); got != want {
			t.Fatalf("twoStrings(%q, %q) = %d, want %d", a, b, got, want)
		}
		if got, want := intervalPalindromes(a), brutePal(a); got != want {
			t.Fatalf("intervalPalindromes(%q) = %d, want %d", a, got, want)
		}
	}
}
