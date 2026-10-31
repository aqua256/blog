package p0394

import (
	"math/rand/v2"
	"strconv"
	"strings"
	"testing"
)

func TestExamples(t *testing.T) {
	cases := map[string]string{
		"3[a]2[bc]":     "aaabcbc",
		"3[a2[c]]":      "accaccacc",
		"2[abc]3[cd]ef": "abcabccdcdcdef",
		"abc3[cd]xyz":   "abccdcdcdxyz",
		"10[a]":         "aaaaaaaaaa",
	}
	for s, want := range cases {
		if got := decodeString(s); got != want {
			t.Errorf("decodeString(%q)=%q want %q", s, got, want)
		}
	}
}

// gen 随机生成一个编码串，同时按定义递归算出它的解码结果
func gen(r *rand.Rand, depth int) (enc, dec string) {
	var e, d strings.Builder
	parts := 1 + r.IntN(3)
	for range parts {
		if depth > 0 && r.IntN(2) == 0 {
			k := 1 + r.IntN(12) // 包含两位数
			ie, id := gen(r, depth-1)
			e.WriteString(strconv.Itoa(k) + "[" + ie + "]")
			d.WriteString(strings.Repeat(id, k))
		} else {
			n := 1 + r.IntN(3)
			for range n {
				c := byte('a' + r.IntN(3))
				e.WriteByte(c)
				d.WriteByte(c)
			}
		}
	}
	return e.String(), d.String()
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(394, 1))
	for range 1000 {
		enc, want := gen(r, 3)
		if got := decodeString(enc); got != want {
			t.Fatalf("decodeString(%q)=%q want %q", enc, got, want)
		}
	}
}
