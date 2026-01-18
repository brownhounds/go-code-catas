package movezeroes

import "testing"

func assertSliceEq(t *testing.T, got, want []int) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("length mismatch: got %d, want %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("at index %d: got %d, want %d; full got=%v want=%v", i, got[i], want[i], got, want)
		}
	}
}

func TestMoveZeroes_Basic(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	MoveZeroes(nums)
	assertSliceEq(t, nums, []int{1, 3, 12, 0, 0})
}

func TestMoveZeroes_NoZeroes(t *testing.T) {
	nums := []int{1, 2, 3}
	MoveZeroes(nums)
	assertSliceEq(t, nums, []int{1, 2, 3})
}

func TestMoveZeroes_AllZeroes(t *testing.T) {
	nums := []int{0, 0, 0}
	MoveZeroes(nums)
	assertSliceEq(t, nums, []int{0, 0, 0})
}

func TestMoveZeroes_AlreadyPartitioned(t *testing.T) {
	nums := []int{1, 2, 3, 0, 0}
	MoveZeroes(nums)
	assertSliceEq(t, nums, []int{1, 2, 3, 0, 0})
}

func TestMoveZeroes_StableOrder(t *testing.T) {
	// relative order of non-zero elements must be preserved
	nums := []int{4, 0, 5, 0, 4, 3, 0, 2}
	MoveZeroes(nums)
	assertSliceEq(t, nums, []int{4, 5, 4, 3, 2, 0, 0, 0})
}

func TestMoveZeroes_SingleElement(t *testing.T) {
	nums := []int{0}
	MoveZeroes(nums)
	assertSliceEq(t, nums, []int{0})

	nums2 := []int{7}
	MoveZeroes(nums2)
	assertSliceEq(t, nums2, []int{7})
}

func TestMoveZeroes_EmptySlice(t *testing.T) {
	var nums []int
	MoveZeroes(nums)
	assertSliceEq(t, nums, nil)
}
