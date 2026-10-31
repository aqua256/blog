package p0279

import "testing"

func TestExamples(t *testing.T) {
	for n, want := range map[int]int{12: 3, 13: 2, 1: 1, 7: 4, 16: 1} {
		if got := numSquares(n); got != want {
			t.Errorf("numSquares(%d) = %d, want %d", n, got, want)
		}
	}
}

// 暴力递归：枚举最后一个平方数
func brute(n int) int {
	if n == 0 {
		return 0
	}
	best := n
	for j := 1; j*j <= n; j++ {
		best = min(best, brute(n-j*j)+1)
	}
	return best
}

func TestBrute(t *testing.T) {
	for n := 1; n <= 28; n++ {
		if got, want := numSquares(n), brute(n); got != want {
			t.Fatalf("numSquares(%d) = %d, want %d", n, got, want)
		}
	}
}

func isSquare(n int) bool {
	for r := 0; r*r <= n; r++ {
		if r*r == n {
			return true
		}
	}
	return false
}

// 数论独立校验：拉格朗日四平方和定理（答案 ≤ 4），
// 勒让德三平方和定理（答案为 4 当且仅当 n = 4^a(8b+7)），答案为 1 当且仅当 n 是平方数
func TestNumberTheory(t *testing.T) {
	for n := 1; n <= 2000; n++ {
		got := numSquares(n)
		m := n
		for m%4 == 0 {
			m /= 4
		}
		switch {
		case isSquare(n) != (got == 1):
			t.Fatalf("numSquares(%d) = %d, square check failed", n, got)
		case (m%8 == 7) != (got == 4):
			t.Fatalf("numSquares(%d) = %d, Legendre check failed", n, got)
		case got < 1 || got > 4:
			t.Fatalf("numSquares(%d) = %d, out of [1,4]", n, got)
		}
	}
}
