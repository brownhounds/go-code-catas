package twosum

// TwoSum returns the indices of the two numbers such that they add up to target.
// If no solution exists, behaviour is undefined for this exercise.
//
// Expected algorithm: Hash Map Lookup (Complement Map)
// Target data structure: map[int]int
func TwoSum(nums []int, target int) (int, int) {
	seen := make(map[int]int)

	for i, value := range nums {
		need := target - value

		if j, ok := seen[need]; ok {
			return j, i
		}

		seen[value] = i
	}

	panic("Pair is not found")
}
