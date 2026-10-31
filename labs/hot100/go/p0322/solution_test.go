package p0322

import (
	"math/rand/v2"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct {
		coins  []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{186, 419, 83, 408}, 6249, 20},
	}
	for _, c := range cases {
		if got := coinChange(c.coins, c.amount); got != c.want {
			t.Errorf("coinChange(%v, %d) = %d, want %d", c.coins, c.amount, got, c.want)
		}
	}
}

// 暴力递归：枚举最后一枚硬币，-1 表示凑不出
func brute(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	best := -1
	for _, c := range coins {
		if c <= amount {
			if r := brute(coins, amount-c); r >= 0 && (best < 0 || r+1 < best) {
				best = r + 1
			}
		}
	}
	return best
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewPCG(322, 2026))
	for range 300 {
		coins := make([]int, 1+rng.IntN(3))
		for i := range coins {
			coins[i] = 2 + rng.IntN(9) // 不含 1，让 -1 情况经常出现
		}
		amount := rng.IntN(25)
		if got, want := coinChange(coins, amount), brute(coins, amount); got != want {
			t.Fatalf("coinChange(%v, %d) = %d, want %d", coins, amount, got, want)
		}
	}
}
