package maximumsubarraysum

// MaxSubarraySum returns the maximum sum of any contiguous subarray.
//
// Constraints:
// - Use O(1) extra space
// - Single pass
//
// Expected algorithm: Kadane's Algorithm (dynamic programming)
// Target concept: iteration with running best/current tracking
func MaxSubarraySum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	current := nums[0]
	best := nums[0]

	for _, v := range nums[1:] {
		current = max(current+v, v)
		best = max(current, best)
	}

	return best
}
