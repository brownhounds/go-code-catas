package rotatesliceright

// RotateRight rotates nums to the right by k steps in-place.
//
// Constraints:
// - Must mutate nums in-place
// - Use O(1) extra space
//
// Expected algorithm: Array Reversal (3-Reverse Trick)
// Target concept: in-place reversal helper
func RotateRight(nums []int, k int) {
	n := len(nums)

	if n == 0 {
		return
	}

	k = k % n
	if k == 0 {
		return
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
