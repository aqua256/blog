package p0062

import "testing"

func TestExamples(t *testing.T) {
	cases := [][3]int{{3, 7, 28}, {3, 2, 3}, {1, 1, 1}, {7, 3, 28}, {3, 3, 6}}
	for _, c := range cases {
		if got := uniquePaths(c[0], c[1]); got != c[2] {
			t.Errorf("uniquePaths(%d, %d) = %d, want %d", c[0], c[1], got, c[2])
		}
	}
}

// 暴力递归：第一步向下或向右
func brute(m, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return brute(m-1, n) + brute(m, n-1)
}

func binom(n, k int) int {
	c := 1
	for x := range k {
		c = c * (n - x) / (x + 1)
	}
	return c
}

func TestBrute(t *testing.T) {
	for m := 1; m <= 10; m++ {
		for n := 1; n <= 10; n++ {
			if got, want := uniquePaths(m, n), brute(m, n); got != want {
				t.Fatalf("uniquePaths(%d, %d) = %d, want %d", m, n, got, want)
			}
		}
	}
	// 组合数校验：C(m+n-2, m-1)
	for m := 1; m <= 17; m++ {
		for n := 1; n <= 17; n++ {
			if got, want := uniquePaths(m, n), binom(m+n-2, m-1); got != want {
				t.Fatalf("uniquePaths(%d, %d) = %d, want C = %d", m, n, got, want)
			}
		}
	}
}
