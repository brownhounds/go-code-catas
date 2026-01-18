package longestuniquesubstring

import "unicode/utf8"

// LongestUniqueSubstring returns the length of the longest substring
// with all unique characters.
//
// Expected algorithm: Sliding Window + Hash Map (Last Seen Index)
// Target DS: map[rune]int, two pointers
func LongestUniqueSubstring(s string) int {
	if utf8.RuneCountInString(s) == 0 {
		return 0
	}

	lastSeen := make(map[rune]int)
	left := 0
	best := 0
	right := 0

	for _, r := range s {
		if prev, ok := lastSeen[r]; ok && prev >= left {
			left = prev + 1
		}

		if right-left+1 > best {
			best = right - left + 1
		}

		lastSeen[r] = right
		right++
	}

	return best
}
