package groupanagrams

import (
	"sort"
	"testing"
)

func normalize(groups [][]string) [][]string {
	out := make([][]string, 0, len(groups))
	for _, g := range groups {
		cp := append([]string(nil), g...)
		sort.Strings(cp)
		out = append(out, cp)
	}
	sort.Slice(out, func(i, j int) bool {
		li := joinWithSep(out[i], "\x00")
		lj := joinWithSep(out[j], "\x00")
		return li < lj
	})
	return out
}

func joinWithSep(xs []string, sep string) string {
	if len(xs) == 0 {
		return ""
	}
	s := xs[0]
	for i := 1; i < len(xs); i++ {
		s += sep + xs[i]
	}
	return s
}

func TestGroupAnagramsUnicode(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  [][]string
	}{
		{
			name:  "empty input",
			words: nil,
			want:  [][]string{},
		},
		{
			name:  "single word",
			words: []string{"żaba"},
			want:  [][]string{{"żaba"}},
		},
		{
			name: "latin with diacritics",
			words: []string{
				"żab", "abż", "bża",
				"ółw", "wół",
			},
			want: [][]string{
				{"żab", "abż", "bża"},
				{"ółw", "wół"},
			},
		},
		{
			name: "greek runes",
			words: []string{
				"αβγ", "βγα", "γαβ",
				"δ", "δ",
			},
			want: [][]string{
				{"αβγ", "βγα", "γαβ"},
				{"δ", "δ"},
			},
		},
		{
			name: "cjk runes",
			words: []string{
				"我爱你", "你我爱", "爱你我",
				"喜欢", "欢喜",
			},
			want: [][]string{
				{"我爱你", "你我爱", "爱你我"},
				{"喜欢", "欢喜"},
			},
		},
		{
			name: "emoji are runes too",
			words: []string{
				"🙂🙃", "🙃🙂",
				"🙂🙂", "🙂🙂",
				"😎",
			},
			want: [][]string{
				{"🙂🙃", "🙃🙂"},
				{"🙂🙂", "🙂🙂"},
				{"😎"},
			},
		},
		{
			name: "mixed unicode and ascii",
			words: []string{
				"aβ", "βa",
				"aβ🙂", "🙂βa", "β🙂a",
				"x",
			},
			want: [][]string{
				{"aβ", "βa"},
				{"aβ🙂", "🙂βa", "β🙂a"},
				{"x"},
			},
		},
		{
			name: "non-anagrams remain separate",
			words: []string{
				"ab", "abc", "abcd", "аб",
			},
			want: [][]string{
				{"ab"},
				{"abc"},
				{"abcd"},
				{"аб"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupAnagrams(tt.words)

			ng := normalize(got)
			nw := normalize(tt.want)

			if len(ng) != len(nw) {
				t.Fatalf("GroupAnagrams(%v) groups=%v, want=%v", tt.words, ng, nw)
			}
			for i := range ng {
				if len(ng[i]) != len(nw[i]) {
					t.Fatalf("GroupAnagrams(%v) group[%d]=%v, want=%v", tt.words, i, ng[i], nw[i])
				}
				for j := range ng[i] {
					if ng[i][j] != nw[i][j] {
						t.Fatalf("GroupAnagrams(%v) group[%d]=%v, want=%v", tt.words, i, ng[i], nw[i])
					}
				}
			}
		})
	}
}
