package lesson7

import (
"testing"
)

func TestSolution(t *testing.T) {
	A := []int{8, 8, 5, 7, 9, 8, 7, 4, 8}
	result := Solution(A)
	if result != 7 {
		t.Errorf("expect 7, but found %v", result)
	}
}
