package tpl_substring

import (
	"slices"
	"testing"
)

func TestCountSubarrays(t *testing.T) {
	if got := countSubarrays([]int{1, -1, 0}, 0); got != 3 {
		t.Fatalf("got %d", got)
	}
}

func TestWindowMax(t *testing.T) {
	if got := windowMax([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3); !slices.Equal(got, []int{3, 3, 5, 5, 6, 7}) {
		t.Fatalf("got %v", got)
	}
}
