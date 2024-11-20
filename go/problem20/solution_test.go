package main

import "testing"

func TestValidParentheses(t *testing.T) {
	// TODO: Add test cases
	// Example:
	// nums := []int{2, 7, 11, 15}
	// target := 9
	// expected := []int{0, 1}
	// result := twoSum(nums, target)
	// if !equal(result, expected) {
	//     t.Errorf("Expected %v, got %v", expected, result)
	// }
}

// Helper function for comparing slices
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
