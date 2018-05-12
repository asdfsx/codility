package lesson6

import (
	"testing"
)

func TestSolution(t *testing.T) {
	A := []int{10, 2, 5, 1, 8, 20}
	result := Solution(A)
	if result != 1 {
		t.Errorf("expect 1, but found %v", result)
	}
}