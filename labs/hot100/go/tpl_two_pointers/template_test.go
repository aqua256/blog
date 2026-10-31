package tpl_two_pointers

import (
	"slices"
	"testing"
)

func TestTwoSumSorted(t *testing.T) {
	if l, r := twoSumSorted([]int{1, 2, 4, 7, 11}, 9); l != 1 || r != 3 {
		t.Fatalf("got %d %d", l, r)
	}
	if l, _ := twoSumSorted([]int{1, 2}, 9); l != -1 {
		t.Fatal("expected miss")
	}
}

func TestFilterInPlace(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	n := filterInPlace(a, func(x int) bool { return x%2 == 0 })
	if !slices.Equal(a[:n], []int{2, 4, 6}) {
		t.Fatalf("got %v", a[:n])
	}
}
