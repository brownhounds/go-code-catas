package slicefunctions

import (
	"reflect"
	"testing"
)

func TestSliceFunctionsSimple(t *testing.T) {
	t.Run("Filter even numbers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		want := []int{2, 4}
		got := Filter(input, func(v int) bool { return v%2 == 0 })
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Filter() = %v, want %v", got, want)
		}
	})

	t.Run("Find first > 3", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		want := 4
		got := Find(input, func(v int) bool { return v > 3 })
		if got != want {
			t.Errorf("Find() = %v, want %v", got, want)
		}
	})

	t.Run("Map squares", func(t *testing.T) {
		input := []int{1, 2, 3}
		want := []int{1, 4, 9}
		got := Map(input, func(v int) int { return v * v })
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Map() = %v, want %v", got, want)
		}
	})

	t.Run("Some > 4?", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		want := true
		got := Some(input, func(v int) bool { return v > 4 })
		if got != want {
			t.Errorf("Some() = %v, want %v", got, want)
		}
	})

	t.Run("Every < 10?", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		want := true
		got := Every(input, func(v int) bool { return v < 10 })
		if got != want {
			t.Errorf("Every() = %v, want %v", got, want)
		}
	})

	t.Run("Includes 3", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		want := true
		got := Includes(input, 3)
		if got != want {
			t.Errorf("Includes() = %v, want %v", got, want)
		}
	})

	t.Run("Includes 7", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		want := false
		got := Includes(input, 7)
		if got != want {
			t.Errorf("Includes() = %v, want %v", got, want)
		}
	})
}
