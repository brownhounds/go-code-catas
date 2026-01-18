package twosum

import "testing"

func TestTwoSum_Basic(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	i, j := TwoSum(nums, target)

	if nums[i]+nums[j] != target {
		t.Fatalf("expected sum %d, got %d", target, nums[i]+nums[j])
	}
}

func TestTwoSum_Negatives(t *testing.T) {
	nums := []int{-3, 4, 1, 2}
	target := -1

	i, j := TwoSum(nums, target)

	if nums[i]+nums[j] != target {
		t.Fatalf("expected sum %d, got %d", target, nums[i]+nums[j])
	}
}

func TestTwoSum_Duplicates(t *testing.T) {
	nums := []int{3, 3}
	target := 6

	i, j := TwoSum(nums, target)

	if i == j {
		t.Fatalf("indices must be different")
	}

	if nums[i]+nums[j] != target {
		t.Fatalf("expected sum %d, got %d", target, nums[i]+nums[j])
	}
}

func TestTwoSum_OrderDoesNotMatter(t *testing.T) {
	nums := []int{1, 5, 3, 7}
	target := 8

	i, j := TwoSum(nums, target)

	if nums[i]+nums[j] != target {
		t.Fatalf("expected sum %d, got %d", target, nums[i]+nums[j])
	}
}
