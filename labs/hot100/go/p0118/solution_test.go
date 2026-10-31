package p0118

import (
	"reflect"
	"testing"
)

func TestExamples(t *testing.T) {
	want := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
	if got := generate(5); !reflect.DeepEqual(got, want) {
		t.Errorf("generate(5) = %v, want %v", got, want)
	}
	if got := generate(1); !reflect.DeepEqual(got, [][]int{{1}}) {
		t.Errorf("generate(1) = %v", got)
	}
}

// 第 i 行第 j 个 = C(i, j)，用乘法公式独立计算
func binom(n, k int) int {
	c := 1
	for x := range k {
		c = c * (n - x) / (x + 1)
	}
	return c
}

func TestBinomial(t *testing.T) {
	for n := 1; n <= 30; n++ {
		got := generate(n)
		if len(got) != n {
			t.Fatalf("generate(%d) has %d rows", n, len(got))
		}
		for i, row := range got {
			if len(row) != i+1 {
				t.Fatalf("generate(%d) row %d has len %d", n, i, len(row))
			}
			for j, v := range row {
				if v != binom(i, j) {
					t.Fatalf("generate(%d)[%d][%d] = %d, want %d", n, i, j, v, binom(i, j))
				}
			}
		}
	}
}
