package validanagram

import "unicode/utf8"

// IsAnagram reports whether s and t are anagrams.
//
// Constraints:
// - Treat input as lower-case a-z only
//
// Expected algorithm: Frequency Count (Hash Map / Fixed Array [26])
// Target data structure: map[rune]int or [26]int
func IsAnagram(s, t string) bool {
	if utf8.RuneCountInString(s) != utf8.RuneCountInString(t) {
		return false
	}

	counts := make(map[rune]int, utf8.RuneCountInString(s))

	for _, r := range s {
		counts[r]++
	}

	for _, r := range t {
		counts[r]--

		if counts[r] < 0 {
			return false
		}
	}

	for _, v := range counts {
		if v != 0 {
			return false
		}
	}

	return true
}
