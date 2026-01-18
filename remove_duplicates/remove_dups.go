package removedup

import "fmt"

// RemoveDuplicates removes duplicates from a sorted slice in-place.
// It returns the new logical length of the slice.
//
// Constraints:
// - Must mutate the input slice
// - Must use O(1) extra space
// - Input slice is sorted
//
// Expected algorithm: Two Pointers (slow / fast)
// Target concepts: slice mutation, index control
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	write := 1

	// @Note first element is always unique by definition
	// start form second position
	for read := 1; read < len(nums); read++ {
		if nums[read] != nums[read-1] {
			nums[write] = nums[read]
			write++
			fmt.Println(nums)
		}
	}

	return write
}
