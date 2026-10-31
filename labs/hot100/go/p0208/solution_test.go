package p0208

import (
	"math/rand/v2"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	tr := Constructor()
	tr.Insert("apple")
	if !tr.Search("apple") || tr.Search("app") || !tr.StartsWith("app") {
		t.Fatal("first block wrong")
	}
	tr.Insert("app")
	if !tr.Search("app") {
		t.Fatal("search app after insert should be true")
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 208))
	word := func() string {
		b := make([]byte, 1+r.IntN(4))
		for i := range b {
			b[i] = "abc"[r.IntN(3)] // 小字母表，前缀碰撞多
		}
		return string(b)
	}
	for range 200 {
		tr := Constructor()
		set := map[string]bool{}
		for op := range 100 {
			w := word()
			switch r.IntN(3) {
			case 0:
				tr.Insert(w)
				set[w] = true
			case 1:
				if got := tr.Search(w); got != set[w] {
					t.Fatalf("op#%d Search(%q) = %v, want %v", op, w, got, set[w])
				}
			case 2:
				want := false
				for s := range set {
					if strings.HasPrefix(s, w) {
						want = true
						break
					}
				}
				if got := tr.StartsWith(w); got != want {
					t.Fatalf("op#%d StartsWith(%q) = %v, want %v", op, w, got, want)
				}
			}
		}
	}
}
