package main

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}
	result := sum(numbers...)
	if result != 6 {
		t.Errorf("result does not match, want 6, have %d", result)
	}
}
