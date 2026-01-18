package rotatesliceright

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

func TestRotateRight_Basic(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	RotateRight(nums, 3)
	assertSliceEq(t, nums, []int{5, 6, 7, 1, 2, 3, 4})
}

func TestRotateRight_KGreaterThanLen(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	RotateRight(nums, 6)
	assertSliceEq(t, nums, []int{3, 4, 1, 2})
}

func TestRotateRight_KEqualsLen(t *testing.T) {
	nums := []int{1, 2, 3}
	RotateRight(nums, 3)
	assertSliceEq(t, nums, []int{1, 2, 3})
}

func TestRotateRight_Zero(t *testing.T) {
	nums := []int{9, 8, 7}
	RotateRight(nums, 0)
	assertSliceEq(t, nums, []int{9, 8, 7})
}

func TestRotateRight_Single(t *testing.T) {
	nums := []int{42}
	RotateRight(nums, 5)
	assertSliceEq(t, nums, []int{42})
}

func TestRotateRight_Empty(t *testing.T) {
	var nums []int
	RotateRight(nums, 2)
	if nums != nil {
		t.Fatalf("expected nil slice, got %v", nums)
	}
}
