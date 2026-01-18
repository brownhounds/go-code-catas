package groupanagrams

import "slices"

// GroupAnagrams groups words by anagram class (Unicode-safe).
//
// Unicode requirement:
// - Treat each word as a sequence of runes (Unicode code points), not bytes.
// - Duplicates matter: "aab" and "aba" are anagrams; "ab" is not.
//
// Task:
//   - Given a slice of strings, group them into buckets where each bucket contains
//     words that are anagrams of each other (by rune multiset).
//
// Rules:
// - Group order does not matter.
// - Word order within a group does not matter.
// - Input strings may contain any Unicode characters.
func GroupAnagrams(words []string) [][]string {
	groups := make(map[string][]string)

	for _, w := range words {
		rs := []rune(w)
		slices.Sort(rs)
		key := string(rs)
		groups[key] = append(groups[key], w)
	}

	out := make([][]string, 0, len(groups))
	for _, g := range groups {
		out = append(out, g)
	}
	return out
}
