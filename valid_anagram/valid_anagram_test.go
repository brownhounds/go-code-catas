package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "basic",
			s:    "listen",
			t:    "silent",
			want: true,
		},
		{
			name: "classic_example",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "length_mismatch",
			s:    "ab",
			t:    "a",
			want: false,
		},
		{
			name: "not_anagram",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "empty_strings",
			s:    "",
			t:    "",
			want: true,
		},
		{
			name: "single_character_true",
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			name: "single_character_false",
			s:    "a",
			t:    "b",
			want: false,
		},
		{
			name: "repeated_characters_true",
			s:    "aabbcc",
			t:    "baccab",
			want: true,
		},
		{
			name: "repeated_characters_false",
			s:    "aabbcc",
			t:    "aabbc",
			want: false,
		},
		{
			name: "same_letters_different_counts",
			s:    "aaab",
			t:    "aabb",
			want: false,
		},
		{
			name: "all_same_letters",
			s:    "zzzz",
			t:    "zzzz",
			want: true,
		},
		{
			name: "early_negative_count",
			s:    "abc",
			t:    "abz",
			want: false,
		},
		{
			name: "false_positive_even_totals_not_anagram",
			s:    "aabb",
			t:    "ccdd",
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsAnagram(tc.s, tc.t)
			if got != tc.want {
				t.Fatalf(
					"IsAnagram(%q, %q) = %v, want %v",
					tc.s, tc.t, got, tc.want,
				)
			}
		})
	}
}
