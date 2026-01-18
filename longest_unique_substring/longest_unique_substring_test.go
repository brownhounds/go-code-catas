package longestuniquesubstring

import "testing"

func TestLongestUniqueSubstring(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{
			name: "empty",
			in:   "",
			want: 0,
		},
		{
			name: "single_char",
			in:   "a",
			want: 1,
		},
		{
			name: "all_unique",
			in:   "abcdef",
			want: 6,
		},
		{
			name: "all_same",
			in:   "bbbbbb",
			want: 1,
		},
		{
			name: "classic_abcabcbb",
			in:   "abcabcbb",
			want: 3, // "abc"
		},
		{
			name: "classic_pwwkew",
			in:   "pwwkew",
			want: 3, // "wke"
		},
		{
			name: "classic_dvdf",
			in:   "dvdf",
			want: 3, // "vdf"
		},
		{
			name: "repeats_within_window",
			in:   "abba",
			want: 2, // "ab" or "ba"
		},
		{
			name: "spaces",
			in:   "a b a",
			want: 3, // "a b" or " b a" depending, longest unique length is 3
		},
		{
			name: "unicode_runes",
			in:   "żółwż",
			want: 4, // "żółw" (must treat as runes, not bytes)
		},
		{
			name: "emoji_runes",
			in:   "🙂🙃🙂",
			want: 2, // "🙂🙃"
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := LongestUniqueSubstring(tt.in)
			if got != tt.want {
				t.Fatalf("LongestUniqueSubstring(%q) = %d, want %d", tt.in, got, tt.want)
			}
		})
	}
}
