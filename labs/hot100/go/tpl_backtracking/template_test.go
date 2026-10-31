package tpl_backtracking

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	if got := fmt.Sprint(combine([]int{1, 2, 3})); got != "[[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]" {
		t.Errorf("combine: got %s", got)
	}
	if got := fmt.Sprint(permute([]int{1, 2, 3})); got != "[[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]" {
		t.Errorf("permute: got %s", got)
	}
}
