package p0049

import (
	"math/rand/v2"
	"slices"
	"strings"
	"testing"
)

// 规范化：组内排序、组间排序，便于比较
func normalize(groups [][]string) []string {
	out := make([]string, 0, len(groups))
	for _, g := range groups {
		g = slices.Clone(g)
		slices.Sort(g)
		out = append(out, strings.Join(g, ","))
	}
	slices.Sort(out)
	return out
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var cnt [26]int
	for i := range a {
		cnt[a[i]-'a']++
		cnt[b[i]-'a']--
	}
	return cnt == [26]int{}
}

// 暴力：逐个放进第一个与之互为异位词的组
func brute(strs []string) [][]string {
	var groups [][]string
	for _, s := range strs {
		placed := false
		for i := range groups {
			if isAnagram(groups[i][0], s) {
				groups[i] = append(groups[i], s)
				placed = true
				break
			}
		}
		if !placed {
			groups = append(groups, []string{s})
		}
	}
	return groups
}

func TestExamples(t *testing.T) {
	cases := []struct {
		in   []string
		want [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}
	for _, c := range cases {
		if got := groupAnagrams(c.in); !slices.Equal(normalize(got), normalize(c.want)) {
			t.Fatalf("in=%v got=%v", c.in, got)
		}
	}
}

func TestRandom(t *testing.T) {
	r := rand.New(rand.NewPCG(49, 49))
	for it := 0; it < 500; it++ {
		strs := make([]string, 1+r.IntN(12))
		for i := range strs {
			b := make([]byte, r.IntN(4))
			for j := range b {
				b[j] = 'a' + byte(r.IntN(3))
			}
			strs[i] = string(b)
		}
		if got, want := normalize(groupAnagrams(strs)), normalize(brute(strs)); !slices.Equal(got, want) {
			t.Fatalf("strs=%v got=%v want=%v", strs, got, want)
		}
	}
}
