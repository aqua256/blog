package p0002

import (
	"math/big"
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := []struct{ a, b, want []int }{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}
	for _, c := range cases {
		if got := toSlice(addTwoNumbers(fromSlice(c.a), fromSlice(c.b))); !slices.Equal(got, c.want) {
			t.Errorf("add(%v,%v) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}

// 逆序数字切片 <-> big.Int
func toBig(d []int) *big.Int {
	x := new(big.Int)
	for i := len(d) - 1; i >= 0; i-- {
		x.Mul(x, big.NewInt(10))
		x.Add(x, big.NewInt(int64(d[i])))
	}
	return x
}

func fromBig(x *big.Int) []int {
	s := x.String()
	out := make([]int, len(s))
	for i := range s {
		out[len(s)-1-i] = int(s[i] - '0')
	}
	return out
}

func randNum(r *rand.Rand) []int {
	n := 1 + r.IntN(25)
	d := make([]int, n)
	for i := range d {
		d[i] = r.IntN(10)
	}
	if n > 1 && d[n-1] == 0 {
		d[n-1] = 1 + r.IntN(9) // 无前导零
	}
	return d
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 2))
	for range 500 {
		a, b := randNum(r), randNum(r)
		want := fromBig(new(big.Int).Add(toBig(a), toBig(b)))
		if got := toSlice(addTwoNumbers(fromSlice(a), fromSlice(b))); !slices.Equal(got, want) {
			t.Fatalf("add(%v,%v) = %v, want %v", a, b, got, want)
		}
	}
}
