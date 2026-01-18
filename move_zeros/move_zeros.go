package movezeroes

// MoveZeroes moves all zeroes to the end of nums in-place,
// while keeping the relative order of non-zero elements.
//
// Constraints:
// - Must mutate nums in-place
// - Must be stable for non-zero elements
// - Use O(1) extra space
//
// Expected algorithm: Two Pointers / Stable Compaction
// Target concepts: in-place slice ops, index control
func MoveZeroes(nums []int) {
	write := 0

	for read := range nums {
		if nums[read] != 0 {
			nums[write] = nums[read]
			write++
		}
	}

	for i := write; i < len(nums); i++ {
		nums[i] = 0
	}
}

func MoveZeroesSwapVersion(nums []int) {
	write := 0

	for read := range nums {
		if nums[read] != 0 {
			nums[read], nums[write] = nums[write], nums[read]
			write++
		}
	}
}
