package removedup

import "testing"

func TestRemoveDuplicates_Basic(t *testing.T) {
	nums := []int{1, 1, 2}
	k := RemoveDuplicates(nums)

	expected := []int{1, 2}
	if k != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), k)
	}

	for i := range k {
		if nums[i] != expected[i] {
			t.Fatalf("at index %d: expected %d, got %d", i, expected[i], nums[i])
		}
	}
}

func TestRemoveDuplicates_MultipleDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := RemoveDuplicates(nums)

	expected := []int{0, 1, 2, 3, 4}
	if k != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), k)
	}

	for i := range k {
		if nums[i] != expected[i] {
			t.Fatalf("at index %d: expected %d, got %d", i, expected[i], nums[i])
		}
	}
}

func TestRemoveDuplicates_AllUnique(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	k := RemoveDuplicates(nums)

	if k != len(nums) {
		t.Fatalf("expected length %d, got %d", len(nums), k)
	}
}

func TestRemoveDuplicates_SingleElement(t *testing.T) {
	nums := []int{1}
	k := RemoveDuplicates(nums)

	if k != 1 {
		t.Fatalf("expected length 1, got %d", k)
	}
}

func TestRemoveDuplicates_EmptySlice(t *testing.T) {
	var nums []int
	k := RemoveDuplicates(nums)

	if k != 0 {
		t.Fatalf("expected length 0, got %d", k)
	}
}
