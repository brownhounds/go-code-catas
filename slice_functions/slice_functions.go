package slicefunctions

import "slices"

// Filter returns a new slice with elements that pass the callback test
func Filter[T any](s []T, callback func(v T) bool) []T {
	out := make([]T, 0, len(s))
	for _, v := range s {
		if callback(v) {
			out = append(out, v)
		}
	}
	return out
}

// Find returns the first element that passes the callback test, or zero value
func Find[T any](s []T, callback func(v T) bool) T {
	for _, v := range s {
		if callback(v) {
			return v
		}
	}
	var zero T
	return zero
}

// Map returns a new slice with the results of applying callback to each element
func Map[T any, R any](s []T, callback func(v T) R) []R {
	out := make([]R, len(s))
	for i, v := range s {
		out[i] = callback(v)
	}
	return out
}

// Some returns true if any element passes the callback test
// @Note as go has idiomic way this is redundant
func Some[T any](s []T, callback func(v T) bool) bool {
	return slices.ContainsFunc(s, callback)
}

// Every returns true if all elements pass the callback test
func Every[T any](s []T, callback func(v T) bool) bool {
	for _, v := range s {
		if !callback(v) {
			return false
		}
	}
	return true
}

// Includes returns true if slice contains the value
// @Note as go has idiomic way this is redundant
func Includes[T comparable](s []T, value T) bool {
	return slices.Contains(s, value)
}
