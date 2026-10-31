package tpl_binary_tree

import (
	"fmt"
	"slices"
	"testing"
)

func TestTemplate(t *testing.T) {
	root := build([]int{1, 2, 3, 4, 5, null, 6, 7})
	if got := postorder(root); got != 5 { // 7-4-2-1-3-6
		t.Errorf("postorder diameter: got %d want 5", got)
	}
	if got := preorder(root); !slices.Equal(got, []int{1, 2, 4, 7}) { // 左视图
		t.Errorf("preorder left view: got %v", got)
	}
	if got := fmt.Sprint(bfs(root)); got != "[[1] [2 3] [4 5 6] [7]]" {
		t.Errorf("bfs: got %s", got)
	}
	if got := bfs(nil); len(got) != 0 {
		t.Errorf("bfs nil: got %v", got)
	}
}
