package p0070

import "testing"

func TestExamples(t *testing.T) {
	for n, want := range map[int]int{1: 1, 2: 2, 3: 3, 5: 8, 45: 1836311903} {
		if got := climbStairs(n); got != want {
			t.Errorf("climbStairs(%d) = %d, want %d", n, got, want)
		}
	}
}

// 暴力递归：最后一步走 1 或 2
func brute(n int) int {
	if n <= 1 {
		return 1
	}
	return brute(n-1) + brute(n-2)
}

func TestBrute(t *testing.T) {
	for n := 1; n <= 25; n++ {
		if got, want := climbStairs(n), brute(n); got != want {
			t.Fatalf("climbStairs(%d) = %d, want %d", n, got, want)
		}
	}
}
