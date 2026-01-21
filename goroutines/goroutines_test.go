package goroutines

import (
	"testing"
)

func TestRunSquares(t *testing.T) {
	got := Run()
	want := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}

	gotMap := make(map[int]bool)
	for _, v := range got {
		gotMap[v] = true
	}

	for _, v := range want {
		if !gotMap[v] {
			t.Errorf("missing expected square: %d", v)
		}
	}

	if len(got) != len(want) {
		t.Errorf("unexpected number of squares: got %d, want %d", len(got), len(want))
	}
}
