package p0155

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestExample(t *testing.T) {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	if got := s.GetMin(); got != -3 {
		t.Fatalf("GetMin=%d want -3", got)
	}
	s.Pop()
	if got := s.Top(); got != 0 {
		t.Fatalf("Top=%d want 0", got)
	}
	if got := s.GetMin(); got != -2 {
		t.Fatalf("GetMin=%d want -2", got)
	}
}

// 对拍：普通切片 + 线性扫描求最小
func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(155, 1))
	for range 200 {
		s := Constructor()
		var ref []int
		for range 100 {
			op := r.IntN(4)
			if len(ref) == 0 {
				op = 0
			}
			switch op {
			case 0:
				v := r.IntN(10) - 5
				s.Push(v)
				ref = append(ref, v)
			case 1:
				s.Pop()
				ref = ref[:len(ref)-1]
			case 2:
				if got, want := s.Top(), ref[len(ref)-1]; got != want {
					t.Fatalf("Top=%d want %d (ref %v)", got, want, ref)
				}
			case 3:
				if got, want := s.GetMin(), slices.Min(ref); got != want {
					t.Fatalf("GetMin=%d want %d (ref %v)", got, want, ref)
				}
			}
		}
	}
}
